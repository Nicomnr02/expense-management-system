package expenseservice

import (
	expensedto "expense-management-system/cmd/expense/dto"
	expensequery "expense-management-system/cmd/expense/repository/query"
	"expense-management-system/dto"
	"expense-management-system/internal/contextkey"
	"expense-management-system/pkg/currency"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (s *expenseServiceImpl) FetchExpense(c *fiber.Ctx, req expensedto.FetchExpenseReq) ([]expensedto.FetchExpenseRes, dto.Page, error) {
	var (
		ctx  = c.Context()
		log  = c.Locals(contextkey.Logger).(*zap.Logger)
		data = []expensedto.FetchExpenseRes{}
		page dto.Page
	)

	query := expensequery.FetchExpense{
		ID:     req.ID,
		Status: req.Status,
		UserID: req.UserID,
		Limit:  req.Page.Limit,
		Offset: req.Page.Page,
	}

	expenses, total, err := s.expenserepository.FetchExpense(ctx, query)
	if err != nil {
		log.Error(err.Error(), zap.Any("query", query))
		return data, page, dto.ErrInternalServer("Fetch expense failed")
	}

	for _, e := range expenses {
		data = append(data,
			expensedto.FetchExpenseRes{
				ID:          e.ID.String(),
				UserID:      e.UserID,
				UserName:    e.UserName,
				AmountIDR:   currency.Rupiah(e.Amount),
				Status:      e.Status,
				SubmittedAt: e.SubmittedAt,
			})
	}

	page = dto.Page{
		Page:  req.Page.Page,
		Limit: req.Page.Limit,
		Total: total,
	}

	return data, page, nil
}
