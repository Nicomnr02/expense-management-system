package mocks

import (
	"context"

	"expense-management-system/database"

	"github.com/stretchr/testify/mock"
)

type TransactionMock struct {
	mock.Mock
}

func (m *TransactionMock) Begin(c context.Context) (database.Tx, error) {
	args := m.Called(c)

	tx, _ := args.Get(0).(database.Tx)
	return tx, args.Error(1)
}

func (m *TransactionMock) Commit(c context.Context, tx database.Tx) error {
	args := m.Called(c, tx)
	return args.Error(0)
}

func (m *TransactionMock) Rollback(c context.Context, tx database.Tx) error {
	args := m.Called(c, tx)
	return args.Error(0)
}
