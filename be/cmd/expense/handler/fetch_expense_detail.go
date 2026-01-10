package expensehandler

import (
	expensedto "expense-management-system/cmd/expense/dto"
	"expense-management-system/model"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *expenseHandlerImpl) FetchExpenseDetail(c *fiber.Ctx) error {
	var request = expensedto.FetchExpenseDetailReq{
		ID: c.Params("id"),
	}

	request.Timestamp = time.Now()
	data, err := h.expenseservice.FetchExpenseDetail(c, request)
	if err != nil {
		return model.Error(c, err, nil)
	}

	return model.Success(c, http.StatusOK, data)
}
