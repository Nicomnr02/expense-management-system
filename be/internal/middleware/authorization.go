package middleware

import (
	"expense-management-system/internal/contextkey"
	"expense-management-system/model"
	"expense-management-system/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

func AuthorizeRole(roles ...string) fiber.Handler {
	roleMap := make(map[string]struct{})
	for _, r := range roles {
		roleMap[r] = struct{}{}
	}

	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals(contextkey.User).(*jwt.AuthClaims)
		if !ok || claims == nil {
			return model.Error(c, model.ErrUnauthorized("Unauthorized"), nil)
		}

		if _, ok := roleMap[claims.Role]; !ok {
			return model.Error(c, model.ErrForbiddenAccess("You are not allowed to access this resource"), nil)
		}
		return c.Next()
	}
}
