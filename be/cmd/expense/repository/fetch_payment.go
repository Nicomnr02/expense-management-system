package expenserepository

import (
	"context"
	expensedomain "expense-management-system/cmd/expense/domain"
	expensequery "expense-management-system/cmd/expense/repository/query"
	"fmt"
	"strings"
)

func (r *expenseRepositoryImpl) FetchPayment(ctx context.Context, q expensequery.FetchPayment) ([]expensedomain.Payment, error) {

	sql := `
	select
		id,
		external_id,
		status,
		created_at,
		updated_at
	from payments p
	`

	var (
		conditions []string
		values     []any
	)

	if len(q.ExternalID) > 0 {
		values = append(values, q.ExternalID)
		conditions = append(conditions, fmt.Sprintf("p.external_id = $%d", len(values)))
	}

	if len(conditions) > 0 {
		sql += " where " + strings.Join(conditions, " and ")
	}

	rows, err := r.DB.Conn.Query(ctx, sql, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []expensedomain.Payment
	for rows.Next() {
		var p expensedomain.Payment
		if err := rows.Scan(
			&p.ID,
			&p.ExternalID,
			&p.Status,
			&p.CreatedAt,
			&p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		payments = append(payments, p)
	}

	return payments, rows.Err()
}
