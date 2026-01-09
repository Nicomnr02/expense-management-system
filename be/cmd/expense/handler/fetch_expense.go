package expensehandler

import (
	expensedto "expense-management-system/cmd/expense/dto"
	"expense-management-system/dto"
	"expense-management-system/internal/contextkey"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (h *expenseHandlerImpl) FetchExpense(c *fiber.Ctx) error {
	var request expensedto.FetchExpenseReq
	var log = c.Locals(contextkey.Logger).(*zap.Logger)

	err := c.BodyParser(&request)
	if err != nil {
		log.Error(err.Error())
		return dto.Error(c, dto.ErrBadRequest("Invalid request data"), nil)
	}

	request.Timestamp = time.Now()
	data, page, err := h.expenseservice.FetchExpense(c, request)
	if err != nil {
		return dto.Error(c, err, nil)
	}

	return dto.SuccessPage(c, data, page)
}
