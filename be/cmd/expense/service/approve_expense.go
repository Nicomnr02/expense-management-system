package expenseservice

import (
	authquery "expense-management-system/cmd/auth/repository/query"
	expensedomain "expense-management-system/cmd/expense/domain"
	expensedto "expense-management-system/cmd/expense/dto"
	expenseenum "expense-management-system/cmd/expense/enum"
	expensequery "expense-management-system/cmd/expense/repository/query"
	"expense-management-system/internal/contextkey"
	"expense-management-system/model"
	"expense-management-system/pkg/jwt"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (s *expenseServiceImpl) ApproveExpense(c *fiber.Ctx, req expensedto.ApprovalReq) (expensedto.ApprovalRes, error) {
	var (
		ctx  = c.Context()
		log  = c.Locals(contextkey.Logger).(*zap.Logger)
		data = expensedto.ApprovalRes{}
	)

	err := s.validator.ValidateStruct(req)
	if err != nil {
		log.Warn(err.Error())
		return data, model.ErrBadRequest(err.Error())
	}

	req.Action = expenseenum.Approve

	expense, approval, payment, err := s.validateApproval(c, req)
	if err != nil {
		return data, err
	}

	expense.ProcessedAt = req.Timestamp
	expense.Status = approval.Status

	tx, err := s.transaction.Begin(ctx)
	defer func() {
		if err != nil {
			_ = s.transaction.Rollback(ctx, tx)
		}
	}()

	err = s.expenserepository.UpdateExpense(ctx, tx, expense)
	if err != nil {
		log.Error(err.Error(), zap.Any("data", approval))
		return data, model.ErrInternalServer("Approve expense failed")
	}

	err = s.expenserepository.CreateApproval(ctx, tx, approval)
	if err != nil {
		log.Error(err.Error(), zap.Any("data", approval))
		return data, model.ErrInternalServer("Approve expense failed")
	}

	if payment != nil {
		err = s.expenserepository.CreatePayment(ctx, tx, *payment)
		if err != nil {
			log.Error(err.Error(), zap.Any("data", payment))
			return data, model.ErrInternalServer("Approve expense failed")
		}
	}

	err = s.transaction.Commit(ctx, tx)
	if err != nil {
		log.Error(err.Error())
		return data, model.ErrInternalServer("Approve expense failed")
	}

	data.Status = approval.Status

	return data, nil
}

func (s *expenseServiceImpl) validateApproval(c *fiber.Ctx, req expensedto.ApprovalReq) (
	expensedomain.Expense,
	expensedomain.Approval,
	*expensedomain.Payment,
	error,
) {

	var (
		ctx = c.Context()
		log = c.Locals(contextkey.Logger).(*zap.Logger)

		expense  expensedomain.Expense
		approval expensedomain.Approval
		payment  *expensedomain.Payment
	)

	query := expensequery.FetchExpense{
		ID: req.ID,
	}

	expenses, _, err := s.expenserepository.FetchExpense(ctx, query)
	if err != nil {
		log.Error(err.Error(), zap.Any("query", query))
		return expense, approval, payment, model.ErrInternalServer("Fetch expense failed")
	}

	if len(expenses) < 1 {
		return expense, approval, payment, model.ErrBadRequest("Expense not found")
	}

	expense = expenses[0]

	if expense.Status != expenseenum.ExpenseAwaitingApproval {
		return expense, approval, payment, model.ErrBadRequest("Expense must be in awaiting approval status")
	}

	claim, ok := c.Locals(contextkey.User).(*jwt.AuthClaims)
	if ok {
		expense.UserID = claim.UserID
	}

	user, err := s.authrepository.FetchUser(ctx,
		authquery.FetchUser{ID: expense.UserID},
	)
	if err != nil {
		log.Warn(err.Error())
		return expense, approval, payment, model.ErrBadRequest("User not found")
	}

	status := expenseenum.ExpenseApproved
	if req.Action == expenseenum.Reject {
		status = expenseenum.ExpenseRejected
	}

	approval = expensedomain.Approval{
		ExpenseID:  expense.ID,
		ApproverID: user.ID,
		Status:     status,
		Notes:      req.Notes,
		CreatedAt:  req.Timestamp,
	}

	if status == expenseenum.ExpenseApproved {
		payment = &expensedomain.Payment{
			ExternalID: expense.ID,
			Status:     expenseenum.ExpensePending,
			CreatedAt:  req.Timestamp,
			UpdatedAt:  req.Timestamp,
		}
	}

	return expense, approval, payment, nil
}
