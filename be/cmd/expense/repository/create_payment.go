package expenserepository

import (
	"context"
	expensedomain "expense-management-system/cmd/expense/domain"
	"expense-management-system/database"
)

func (r *expenseRepositoryImpl) CreatePayment(ctx context.Context, tx database.Tx, data expensedomain.Payment) error {
	query := `
		insert into payments (
			external_id,
			status,
			created_at,
			updated_at
		)
		values ($1, $2, $3, $4)
	`

	_, err := tx.Exec(
		ctx,
		query,
		data.ExternalID,
		data.Status,
		data.CreatedAt,
		data.UpdatedAt,
	)

	return err
}
