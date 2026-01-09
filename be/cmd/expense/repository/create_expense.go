package expenserepository

import (
	"context"
	expensedomain "expense-management-system/cmd/expense/domain"
	"expense-management-system/database"

	"github.com/google/uuid"
)

func (r *expenseRepositoryImpl) CreateExpense(c context.Context, tx database.Tx, data expensedomain.Expense) error {
	query := `
		INSERT INTO expenses (id, user_id, amount, description, receipt_url, status, submitted_at, processed_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	if data.ID == uuid.Nil {
		data.ID = uuid.New()
	}

	_, err := tx.Exec(c, query,
		data.ID,
		data.UserID,
		data.Amount,
		data.Description,
		data.ReceiptURL,
		data.Status,
		data.SubmittedAt,
		data.ProcessedAt,
	)
	return err
}
