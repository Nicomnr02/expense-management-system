package authhandler

import (
	authservice "expense-management-system/cmd/auth/service"
	"expense-management-system/internal/middleware"
	"expense-management-system/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

type authHandlerImpl struct {
	server      fiber.Router
	authservice authservice.AuthService
	JWTManager  *jwt.JWTManager
}

func New(router fiber.Router, authservice authservice.AuthService, JWTManager *jwt.JWTManager) {
	handler := authHandlerImpl{
		server:      router,
		authservice: authservice,
		JWTManager:  JWTManager,
	}

	auth := router.Group("/auth")
	auth.Post("/login", handler.Login)
	auth.Get("/read-token", middleware.Authenticate(JWTManager), handler.ReadToken)
}
