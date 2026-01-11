package expenserepository

import (
	"context"
	expensedomain "expense-management-system/cmd/expense/domain"
	"expense-management-system/database"
	"time"
)

func (r *expenseRepositoryImpl) UpdatePayment(c context.Context, tx database.Tx, data expensedomain.Payment) error {
	query := `update payments set status = $2, updated_at = $3
		where external_id = $1`

	_, err := tx.Exec(c, query,
		data.ExternalID,
		data.Status,
		time.Now(),
	)
	return err
}
