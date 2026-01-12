package authhandler

import (
	"expense-management-system/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *authHandlerImpl) ReadToken(c *fiber.Ctx) error {

	data, err := h.authservice.ReadToken(c)
	if err != nil {
		return model.Error(c, err, nil)
	}

	return model.Success(c, http.StatusOK, data)
}
