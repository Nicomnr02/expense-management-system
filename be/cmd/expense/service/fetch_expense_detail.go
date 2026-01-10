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

	expenses, _, err := s.expenserepository.FetchExpense(ctx,
		expensequery.FetchExpense{
			ID: req.ID,
		})

	if err != nil {
		log.Error(err.Error(), zap.Any("expense_id", req.ID))
		return data, model.ErrInternalServer("Fetch expense detail failed")
	}

	if len(expenses) < 1 {
		return data, model.ErrBadRequest("Expense not found")
	}

	approvals, err := s.expenserepository.FetchApproval(ctx,
		expensequery.FetchApproval{
			ExpenseID: req.ID,
		})

	if err != nil {
		log.Error(err.Error(), zap.Any("expense_id", req.ID))
		return data, model.ErrInternalServer("Fetch expense detail failed")
	}

	payments, err := s.expenserepository.FetchPayment(ctx,
		expensequery.FetchPayment{
			ExternalID: req.ID,
		})

	if err != nil {
		log.Error(err.Error(), zap.Any("expense_id", req.ID))
		return data, model.ErrInternalServer("Fetch expense detail failed")
	}

	expense := expenses[0]
	data.Expense = expensedto.FetchExpenseDetailExpenseRes{
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

	if len(approvals) > 0 {
		approval := approvals[0]
		data.Appproval = &expensedto.FetchExpenseDetailApprovalRes{
			ID:           approval.ID,
			ApproverID:   approval.ApproverID,
			ApproverName: approval.ApproverName,
			ApproverRole: approval.ApproverRole,
			Status:       approval.Status,
			Notes:        approval.Notes,
			CreatedAt:    approval.CreatedAt,
		}
	}

	if len(payments) > 0 {
		payment := payments[0]
		data.Payment = &expensedto.FetchExpenseDetailPaymentRes{
			ID:         payment.ID,
			ExternalID: payment.ExternalID,
			Status:     payment.Status,
			CreatedAt:  payment.CreatedAt,
			UpdatedAt:  payment.UpdatedAt,
		}
	}

	return data, nil
}
