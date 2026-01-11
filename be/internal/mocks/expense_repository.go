package mocks

import (
	"context"

	expensedomain "expense-management-system/cmd/expense/domain"
	expensequery "expense-management-system/cmd/expense/repository/query"
	"expense-management-system/database"

	"github.com/stretchr/testify/mock"
)

type ExpenseRepositoryMock struct {
	mock.Mock
}

func (m *ExpenseRepositoryMock) CreateExpense(
	ctx context.Context,
	tx database.Tx,
	data expensedomain.Expense,
) error {
	args := m.Called(ctx, tx, data)
	return args.Error(0)
}

func (m *ExpenseRepositoryMock) CreateApproval(
	ctx context.Context,
	tx database.Tx,
	data expensedomain.Approval,
) error {
	args := m.Called(ctx, tx, data)
	return args.Error(0)
}

func (m *ExpenseRepositoryMock) CreatePayment(
	ctx context.Context,
	tx database.Tx,
	data expensedomain.Payment,
) error {
	args := m.Called(ctx, tx, data)
	return args.Error(0)
}

func (m *ExpenseRepositoryMock) UpdateExpense(
	ctx context.Context,
	tx database.Tx,
	data expensedomain.Expense,
) error {
	args := m.Called(ctx, tx, data)
	return args.Error(0)
}

func (m *ExpenseRepositoryMock) UpdatePayment(
	ctx context.Context,
	tx database.Tx,
	data expensedomain.Payment,
) error {
	args := m.Called(ctx, tx, data)
	return args.Error(0)
}

func (m *ExpenseRepositoryMock) FetchExpense(
	ctx context.Context,
	q expensequery.FetchExpense,
) ([]expensedomain.Expense, int, error) {
	args := m.Called(ctx, q)

	data, _ := args.Get(0).([]expensedomain.Expense)
	total, _ := args.Get(1).(int)

	return data, total, args.Error(2)
}

func (m *ExpenseRepositoryMock) FetchApproval(
	ctx context.Context,
	q expensequery.FetchApproval,
) ([]expensedomain.Approval, error) {
	args := m.Called(ctx, q)

	data, _ := args.Get(0).([]expensedomain.Approval)
	return data, args.Error(1)
}

func (m *ExpenseRepositoryMock) FetchPayment(
	ctx context.Context,
	q expensequery.FetchPayment,
) ([]expensedomain.Payment, error) {
	args := m.Called(ctx, q)

	data, _ := args.Get(0).([]expensedomain.Payment)
	return data, args.Error(1)
}
