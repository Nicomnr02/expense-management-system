package authhandler

import (
	authservice "expense-management-system/cmd/auth/service"

	"github.com/gofiber/fiber/v2"
)

type authHandlerImpl struct {
	server      *fiber.App
	authservice authservice.AuthService
}

func New(server *fiber.App, authservice authservice.AuthService) {
	handler := authHandlerImpl{
		server:      server,
		authservice: authservice,
	}

	auth := server.Group("auth")
	auth.Post("/login", handler.Login)
}
