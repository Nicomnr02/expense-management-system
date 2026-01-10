package expenserepository

import (
	"context"
	expensedomain "expense-management-system/cmd/expense/domain"
	expensequery "expense-management-system/cmd/expense/repository/query"
	"fmt"
	"strings"
)

func (r *expenseRepositoryImpl) FetchApproval(ctx context.Context, q expensequery.FetchApproval) ([]expensedomain.Approval, error) {

	sql := `
	select
		a.id,
		a.expense_id,
		a.approver_id,
		u.name,
		u.role,
		a.status,
		a.notes,
		a.created_at
	from approvals a
	join users u on a.approver_id = u.id
	`

	var (
		conditions []string
		values     []any
	)

	if len(q.ExpenseID) > 0 {
		values = append(values, q.ExpenseID)
		conditions = append(conditions, fmt.Sprintf("a.expense_id = $%d", len(values)))
	}

	if len(conditions) > 0 {
		sql += " where " + strings.Join(conditions, " and ")
	}

	rows, err := r.DB.Conn.Query(ctx, sql, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var approval []expensedomain.Approval
	for rows.Next() {
		var a expensedomain.Approval
		if err := rows.Scan(
			&a.ID,
			&a.ExpenseID,
			&a.ApproverID,
			&a.ApproverName,
			&a.ApproverRole,
			&a.Status,
			&a.Notes,
			&a.CreatedAt,
		); err != nil {
			return nil, err
		}
		approval = append(approval, a)
	}

	return approval, rows.Err()
}
