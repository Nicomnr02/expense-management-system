package expensehandler

import (
	expensedto "expense-management-system/cmd/expense/dto"
	"expense-management-system/model"
	"expense-management-system/internal/contextkey"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (h *expenseHandlerImpl) CreateExpense(c *fiber.Ctx) error {
	var request expensedto.CreateExpenseReq
	var log = c.Locals(contextkey.Logger).(*zap.Logger)

	err := c.BodyParser(&request)
	if err != nil {
		log.Error(err.Error())
		return model.Error(c, model.ErrBadRequest("Invalid request data"), nil)
	}

	request.Timestamp = time.Now()
	data, err := h.expenseservice.CreateExpense(c, request)
	if err != nil {
		return model.Error(c, err, nil)
	}

	return model.Success(c, http.StatusOK, data)

}
