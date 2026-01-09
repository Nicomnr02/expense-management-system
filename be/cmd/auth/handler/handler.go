package authhandler

import (
	authservice "expense-management-system/cmd/auth/service"
	"expense-management-system/pkg/httpserver"
)

type authHandlerImpl struct {
	server      *httpserver.Server
	authservice authservice.AuthService
}

func New(server *httpserver.Server, authservice authservice.AuthService) {
	handler := authHandlerImpl{
		server:      server,
		authservice: authservice,
	}

	auth := server.Group("auth")
	auth.Post("/login", server.Use(handler.Login))
}
