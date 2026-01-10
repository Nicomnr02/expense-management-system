package expensehandler

import (
	expensedto "expense-management-system/cmd/expense/dto"
	"expense-management-system/internal/contextkey"
	"expense-management-system/model"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (h *expenseHandlerImpl) RejectExpense(c *fiber.Ctx) error {
	var request expensedto.ApprovalReq
	var log = c.Locals(contextkey.Logger).(*zap.Logger)

	err := c.BodyParser(&request)
	if err != nil {
		log.Error(err.Error())
		return model.Error(c, model.ErrBadRequest("Invalid request data"), nil)
	}

	request.ID = c.Params("id")
	request.Timestamp = time.Now()

	data, err := h.expenseservice.RejectExpense(c, request)
	if err != nil {
		return model.Error(c, err, nil)
	}

	return model.Success(c, http.StatusOK, data)
}
