package expenseservice

import (
	expensedto "expense-management-system/cmd/expense/dto"
	expenseenum "expense-management-system/cmd/expense/enum"
	"expense-management-system/internal/contextkey"
	"expense-management-system/model"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (s *expenseServiceImpl) RejectExpense(c *fiber.Ctx, req expensedto.ApprovalReq) (expensedto.ApprovalRes, error) {
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

	req.Action = expenseenum.Reject

	expense, approval, _, err := s.validateApproval(c, req)
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
		return data, model.ErrInternalServer("Reject expense failed")
	}

	err = s.transaction.Commit(ctx, tx)
	if err != nil {
		log.Error(err.Error())
		return data, model.ErrInternalServer("Reject expense failed")
	}

	data.Status = approval.Status

	return data, nil
}
