package expenseservice

import (
	expensedto "expense-management-system/cmd/expense/dto"
	expensequery "expense-management-system/cmd/expense/repository/query"
	"expense-management-system/internal/contextkey"
	middlewareenum "expense-management-system/internal/middleware/enum"
	"expense-management-system/model"
	"expense-management-system/pkg/currency"
	"expense-management-system/pkg/jwt"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (s *expenseServiceImpl) FetchExpense(c *fiber.Ctx, req expensedto.FetchExpenseReq) ([]expensedto.FetchExpenseRes, model.Pagination, error) {
	var (
		ctx   = c.Context()
		log   = c.Locals(contextkey.Logger).(*zap.Logger)
		claim = c.Locals(contextkey.User).(*jwt.AuthClaims)
		data  = []expensedto.FetchExpenseRes{}
		page  model.Pagination
	)

	query := expensequery.FetchExpense{
		ID:         req.ID,
		Status:     req.Status,
		Pagination: req.Pagination,
	}

	if claim.Role == middlewareenum.EMPLOYEEROLE {
		query.UserID = claim.UserID
	}

	expenses, total, err := s.expenserepository.FetchExpense(ctx, query)
	if err != nil {
		log.Error(err.Error(), zap.Any("query", query))
		return data, page, model.ErrInternalServer("Fetch expense failed")
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

	page = req.Pagination
	page.Total = total

	return data, page, nil
}
