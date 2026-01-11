package expenseservice

import (
	"encoding/json"
	authquery "expense-management-system/cmd/auth/repository/query"
	expensedomain "expense-management-system/cmd/expense/domain"
	expensedto "expense-management-system/cmd/expense/dto"
	expenseenum "expense-management-system/cmd/expense/enum"
	"expense-management-system/internal/contextkey"
	"expense-management-system/internal/job"
	"expense-management-system/model"
	"expense-management-system/pkg/jwt"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (s *expenseServiceImpl) CreateExpense(c *fiber.Ctx, req expensedto.CreateExpenseReq) (expensedto.CreateExpenseRes, error) {
	var (
		ctx  = c.Context()
		log  = c.Locals(contextkey.Logger).(*zap.Logger)
		data expensedto.CreateExpenseRes
	)

	err := s.validator.ValidateStruct(req)
	if err != nil {
		log.Warn(err.Error())
		return data, model.ErrBadRequest(err.Error())
	}

	if req.AmountIDR < s.cfg.MinExpenseAmount || req.AmountIDR > s.cfg.MaxExpenseAmount {
		return data, model.ErrBadRequest(
			fmt.Sprintf("Amount must be between %d and %d",
				s.cfg.MinExpenseAmount,
				s.cfg.MaxExpenseAmount,
			))
	}

	expense := expensedomain.Expense{
		ID:          uuid.New(),
		Amount:      req.AmountIDR,
		Description: req.Description,
		ReceiptURL:  req.ReceiptURL,
		SubmittedAt: req.Timestamp,
		ProcessedAt: req.Timestamp,
	}

	expense.Status = expenseenum.ExpenseAwaitingApproval

	if req.AmountIDR < s.cfg.ApprovalThreshold {
		expense.Status = expenseenum.ExpenseAutoApproved
	}

	var (
		approval *expensedomain.Approval
		payment  *expensedomain.Payment
		payTask  *job.Task
	)

	if expense.Status == expenseenum.ExpenseAutoApproved {
		approval = &expensedomain.Approval{
			ExpenseID:  expense.ID,
			ApproverID: 1,
			Status:     expense.Status,
			Notes:      "Auto-approved by system.",
			CreatedAt:  req.Timestamp,
		}

		payment = &expensedomain.Payment{
			ExternalID: expense.ID,
			Status:     expenseenum.ExpensePending,
			CreatedAt:  req.Timestamp,
			UpdatedAt:  req.Timestamp,
		}

		paymentTask := expensedto.PaymentReq{
			ExternalID: expense.ID.String(),
			Amount:     expense.Amount,
		}

		paymentbyte, err := json.Marshal(&paymentTask)
		if err != nil {
			log.Error(err.Error(), zap.Any("payment", payment))
			return data, model.ErrInternalServer("Create expense failed")
		}

		payTask = &job.Task{
			Action:  expenseenum.Pay,
			Payload: paymentbyte,
		}

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
		return data, model.ErrBadRequest("User not found")
	}

	tx, err := s.transaction.Begin(ctx)
	defer func() {
		if err != nil {
			_ = s.transaction.Rollback(ctx, tx)
		}
	}()

	err = s.expenserepository.CreateExpense(ctx, tx, expense)
	if err != nil {
		log.Error(err.Error(), zap.Any("data", expense))
		return data, model.ErrInternalServer("Create expense failed")
	}

	if approval != nil {
		err = s.expenserepository.CreateApproval(ctx, tx, *approval)
		if err != nil {
			log.Error(err.Error(), zap.Any("data", approval))
			return data, model.ErrInternalServer("Create expense failed")
		}
	}

	if payment != nil {
		err = s.expenserepository.CreatePayment(ctx, tx, *payment)
		if err != nil {
			log.Error(err.Error(), zap.Any("data", payment))
			return data, model.ErrInternalServer("Create expense failed")
		}
	}

	if payTask != nil {
		log.Info("sending paytask queue...")
		err := s.jobClient.Enqueue(*payTask)
		if err != nil {
			log.Error(err.Error(), zap.Any("data", *payTask))
			return data, model.ErrInternalServer("Create expense failed")
		}
	}

	err = s.transaction.Commit(ctx, tx)
	if err != nil {
		log.Error(err.Error())
		return data, model.ErrInternalServer("Create expense failed")
	}

	data = expensedto.CreateExpenseRes{
		ID:          expense.ID,
		UserID:      user.ID,
		UserName:    user.Name,
		AmountIDR:   expense.Amount,
		SubmittedAt: expense.SubmittedAt,
		ProcessedAt: expense.ProcessedAt,
		Status:      expense.Status,
	}

	return data, nil

}
