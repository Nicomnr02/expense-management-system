package authhandler

import (
	authservice "expense-management-system/cmd/auth/service"
	"expense-management-system/pkg/httpserver"

	"go.uber.org/zap"
)

type authHandlerImpl struct {
	server      httpserver.Server
	authservice authservice.AuthService
	logger      *zap.Logger
}

func New(server httpserver.Server, authservice authservice.AuthService, logger *zap.Logger) {
	handler := authHandlerImpl{
		server:      server,
		authservice: authservice,
		logger:      logger,
	}

	auth := server.Group("auth")
	auth.Post("/login", server.Use(handler.Login))
}
