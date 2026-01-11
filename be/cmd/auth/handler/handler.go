package authhandler

import (
	authservice "expense-management-system/cmd/auth/service"

	"github.com/gofiber/fiber/v2"
)

type authHandlerImpl struct {
	server      fiber.Router
	authservice authservice.AuthService
}

func New(router fiber.Router, authservice authservice.AuthService) {
	handler := authHandlerImpl{
		server:      router,
		authservice: authservice,
	}

	auth := router.Group("/auth")
	auth.Post("/login", handler.Login)
}
