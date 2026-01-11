package expenseservice

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	expensedto "expense-management-system/cmd/expense/dto"
	expenseenum "expense-management-system/cmd/expense/enum"
	expensequery "expense-management-system/cmd/expense/repository/query"
	"expense-management-system/internal/contextkey"
	"expense-management-system/internal/job"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func (s *expenseServiceImpl) PayExpense(c context.Context, task job.Task) error {

	var (
		log        = c.Value(contextkey.Worker).(*zap.Logger)
		paymentReq = expensedto.PaymentReq{}
	)

	log.Info("received paytask queue...",
		zap.Int("retried", task.RetryCount),
		zap.Int("max_retry", task.MaxRetry),
	)

	err := json.Unmarshal(task.Payload, &paymentReq)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	expenses, _, err := s.expenserepository.FetchExpense(c,
		expensequery.FetchExpense{ID: paymentReq.ExternalID})

	payments, err := s.expenserepository.FetchPayment(c,
		expensequery.FetchPayment{ExternalID: paymentReq.ExternalID})

	if len(expenses) < 1 || len(payments) < 1 {
		log.Error(err.Error())
		return errors.New("data not found")
	}

	expense := expenses[0]
	payment := payments[0]

	if expense.Status == expenseenum.ExpenseCompleted {
		return nil
	}

	response, err := s.doPayment(c, task)

	failed := err != nil || response.Data.Status != expenseenum.PaymentSuccess

	if failed && (task.RetryCount >= task.MaxRetry) {
		expense.Status = expenseenum.ExpenseUnderReview
		payment.Status = expenseenum.ExpenseUnderReview

	} else if failed {
		return fmt.Errorf("Failed payment: %s", response.Message)

	} else {
		expense.Status = expenseenum.ExpenseCompleted
		payment.Status = expenseenum.PaymentSuccess
	}

	tx, err := s.transaction.Begin(c)
	defer func() {
		if err != nil {
			_ = s.transaction.Rollback(c, tx)
		}
	}()

	err = s.expenserepository.UpdateExpense(c, tx, expense)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	err = s.expenserepository.UpdatePayment(c, tx, payment)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	err = s.transaction.Commit(c, tx)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	return nil
}

func (s *expenseServiceImpl) doPayment(c context.Context, task job.Task) (expensedto.PaymentRes, error) {
	var (
		res expensedto.PaymentRes
		log = c.Value(contextkey.Worker).(*zap.Logger)
	)

	req, err := http.NewRequestWithContext(
		c,
		http.MethodPost,
		s.cfg.PaymentURL,
		bytes.NewBuffer(task.Payload),
	)
	if err != nil {
		log.Error(err.Error())
		return res, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err.Error())
		return res, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		log.Error(err.Error())
		return res, err
	}

	return res, nil
}
