package authhandler

import (
	authdto "expense-management-system/cmd/auth/dto"
	"expense-management-system/internal/contextkey"
	"expense-management-system/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (h *authHandlerImpl) Login(c *fiber.Ctx) error {

	var request authdto.LoginReq
	var log = c.Locals(contextkey.Logger).(*zap.Logger)

	err := c.BodyParser(&request)
	if err != nil {
		log.Error(err.Error())
		return model.Error(c, model.ErrBadRequest("Invalid request data"), nil)
	}

	data, err := h.authservice.Login(c, request)
	if err != nil {
		return model.Error(c, err, nil)
	}

	return model.Success(c, http.StatusOK, data)
}
