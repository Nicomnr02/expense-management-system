package expenserepository

import (
	"context"
	expensedomain "expense-management-system/cmd/expense/domain"
	"expense-management-system/database"
)

type ExpenseRepository interface {
	CreateExpense(ctx context.Context, tx database.Tx, data expensedomain.Expense) error
	CreateApproval(c context.Context, tx database.Tx, data expensedomain.Approval) error
	CreatePayment(ctx context.Context, tx database.Tx, data expensedomain.Payment) error
}

type expenseRepositoryImpl struct {
	DB *database.Database
}

func New(DB *database.Database) ExpenseRepository {
	return &expenseRepositoryImpl{DB: DB}
}
