package expenseservice

import (
	expensedto "expense-management-system/cmd/expense/dto"
	expensequery "expense-management-system/cmd/expense/repository/query"
	"expense-management-system/internal/contextkey"
	"expense-management-system/model"
	"expense-management-system/pkg/currency"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (s *expenseServiceImpl) FetchExpenseDetail(c *fiber.Ctx, req expensedto.FetchExpenseDetailReq) (expensedto.FetchExpenseDetailRes, error) {
	var (
		ctx  = c.Context()
		log  = c.Locals(contextkey.Logger).(*zap.Logger)
		data = expensedto.FetchExpenseDetailRes{}
	)

	query := expensequery.FetchExpense{
		ID: req.ID,
	}

	expenses, _, err := s.expenserepository.FetchExpense(ctx, query)
	if err != nil {
		log.Error(err.Error(), zap.Any("query", query))
		return data, model.ErrInternalServer("Fetch expense failed")
	}

	if len(expenses) < 1 {
		return data, model.ErrBadRequest("Expense not found")
	}

	expense := expenses[0]

	data = expensedto.FetchExpenseDetailRes{
		ID:          expense.ID,
		UserID:      expense.UserID,
		UserName:    expense.UserName,
		AmountIDR:   currency.Rupiah(expense.Amount),
		Description: expense.Description,
		ReceiptURL:  expense.ReceiptURL,
		Status:      expense.Status,
		SubmittedAt: expense.SubmittedAt,
		ProcessedAt: expense.ProcessedAt,
	}

	return data, nil
}
