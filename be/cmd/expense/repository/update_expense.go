package expenserepository

import (
	"context"
	expensedomain "expense-management-system/cmd/expense/domain"
	"expense-management-system/database"

	"github.com/google/uuid"
)

func (r *expenseRepositoryImpl) UpdateExpense(c context.Context, tx database.Tx, data expensedomain.Expense) error {
	query := `
	UPDATE expenses
	SET
		user_id = $2,
		amount = $3,
		description = $4,
		receipt_url = $5,
		status = $6,
		submitted_at = $7,
		processed_at = $8
	WHERE id = $1	
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
