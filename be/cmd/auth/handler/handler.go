package authhandler

import (
	authservice "expense-management-system/cmd/auth/service"
	"expense-management-system/dto"
	"expense-management-system/pkg/httpserver"
	"net/http"
)

type authHandlerImpl struct {
	server      httpserver.Server
	authservice authservice.AuthService
}

func (h *authHandlerImpl) Login(ctx httpserver.Context) error {
	return dto.Success(ctx, http.StatusOK, "mantap")
}

func New(server httpserver.Server, authservice authservice.AuthService) {
	// handler := authHandlerImpl{
	// 	server:      server,
	// 	authservice: authservice,
	// }

	// auth := server.Group("auth")
}
