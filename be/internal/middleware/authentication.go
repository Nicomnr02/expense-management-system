package middleware

import (
	"strings"

	"expense-management-system/internal/contextkey"
	"expense-management-system/model"
	"expense-management-system/pkg/jwt"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Authenticate(manager *jwt.JWTManager) fiber.Handler {

	return func(c *fiber.Ctx) error {
		log := c.Locals(contextkey.Logger).(*zap.Logger)

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			log.Warn("missing Authorization header")
			return model.Error(c, model.ErrUnauthorized("Unauthorized"), nil)
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			log.Warn("invalid Authorization header format", zap.String("header", authHeader))
			return model.Error(c, model.ErrUnauthorized("Unauthorized"), nil)
		}

		tokenString := parts[1]

		claims, err := manager.ValidateAccessToken(tokenString)
		if err != nil {
			log.Warn("invalid access token", zap.Error(err))
			return model.Error(c, model.ErrUnauthorized("Unauthorized"), nil)
		}

		c.Locals(contextkey.User, claims)

		return c.Next()
	}
}
