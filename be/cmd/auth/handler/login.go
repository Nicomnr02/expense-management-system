package authhandler

import (
	authdto "expense-management-system/cmd/auth/dto"
	"expense-management-system/dto"
	"expense-management-system/pkg/httpserver"
	"net/http"
)

func (h *authHandlerImpl) Login(ctx httpserver.Context) error {

	var request authdto.LoginReq
	err := h.server.Parse(ctx, &request)
	if err != nil {
		h.logger.Error(err.Error())
		return dto.Error(ctx, dto.ErrBadRequest("Invalid request data"), nil)
	}

	c := h.server.Context(ctx)
	data, err := h.authservice.Login(c, request)
	if err != nil {
		return dto.Error(ctx, err, nil)
	}

	return dto.Success(ctx, http.StatusOK, data)
}
