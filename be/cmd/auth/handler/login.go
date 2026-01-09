package authhandler

import (
	authdto "expense-management-system/cmd/auth/dto"
	"expense-management-system/dto"
	"expense-management-system/pkg/httpserver"
	"net/http"
)

func (h *authHandlerImpl) Login(ctx httpserver.Context) error {

	var request authdto.LoginReq
	var log = httpserver.UseLogger(ctx)

	err := h.server.Parse(ctx, &request)
	if err != nil {
		log.Error(err.Error())
		return dto.Error(ctx, dto.ErrBadRequest("Invalid request data"), nil)
	}

	data, err := h.authservice.Login(ctx, request)
	if err != nil {
		return dto.Error(ctx, err, nil)
	}

	return dto.Success(ctx, http.StatusOK, data)
}
