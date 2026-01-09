package expenserepository

import (
	"context"
	expensedomain "expense-management-system/cmd/expense/domain"
	"expense-management-system/database"
)

func (r *expenseRepositoryImpl) CreatePayment(ctx context.Context, tx database.Tx, data expensedomain.Payment) error {
	query := `
		INSERT INTO payments (
			external_id,
			status,
			retry_attempts,
			created_at,
			updated_at
		)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := tx.Exec(
		ctx,
		query,
		data.ExternalID,
		data.Status,
		data.RetryAttempts,
		data.CreatedAt,
		data.UpdatedAt,
	)

	return err
}
