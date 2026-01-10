package expenserepository

import (
	"context"
	expensedomain "expense-management-system/cmd/expense/domain"
	expensequery "expense-management-system/cmd/expense/repository/query"
	"expense-management-system/pkg/pagination"
	"fmt"
	"strings"
)

func (r *expenseRepositoryImpl) FetchExpense(ctx context.Context, q expensequery.FetchExpense) ([]expensedomain.Expense, int, error) {

	whereSQL, values := r.queryFetchExpense(q)

	sql := `
		select
			e.id,
			e.user_id,
			u.name,
			e.amount,
			e.description,
			e.receipt_url,
			e.status,
			e.submitted_at,
			e.processed_at
		from expenses e 
		join users u on u.id = e.user_id
	` + whereSQL + `
		order by submitted_at desc
	`

	fetchValues := append([]any{}, values...)

	sql, fetchValues = r.paginateFetchExpense(sql, fetchValues, q)

	rows, err := r.DB.Conn.Query(ctx, sql, fetchValues...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var expenses []expensedomain.Expense
	for rows.Next() {
		var e expensedomain.Expense
		if err := rows.Scan(
			&e.ID,
			&e.UserID,
			&e.UserName,
			&e.Amount,
			&e.Description,
			&e.ReceiptURL,
			&e.Status,
			&e.SubmittedAt,
			&e.ProcessedAt,
		); err != nil {
			return nil, 0, err
		}
		expenses = append(expenses, e)
	}

	total, err := r.countFetchExpense(ctx, q)
	if err != nil {
		return nil, 0, err
	}

	return expenses, total, rows.Err()
}

func (r *expenseRepositoryImpl) countFetchExpense(ctx context.Context, q expensequery.FetchExpense) (int, error) {
	whereSQL, values := r.queryFetchExpense(q)

	countSQL := `select count(e.id) from expenses e join users u on u.id = e.user_id` + whereSQL

	var total int
	if err := r.DB.Conn.QueryRow(ctx, countSQL, values...).Scan(&total); err != nil {
		return 0, err
	}

	return total, nil
}

func (r *expenseRepositoryImpl) paginateFetchExpense(sql string, values []any, q expensequery.FetchExpense) (string, []any) {
	pagination := pagination.New(q.Page, q.Limit)

	sql += fmt.Sprintf(" limit $%d", len(values)+1)
	values = append(values, pagination.Limit)

	sql += fmt.Sprintf(" offset $%d", len(values)+1)
	values = append(values, pagination.Offset)

	return sql, values
}

func (r *expenseRepositoryImpl) queryFetchExpense(q expensequery.FetchExpense) (string, []any) {
	var (
		conditions []string
		values     []any
	)

	if q.ID != "" {
		values = append(values, q.ID)
		conditions = append(conditions, fmt.Sprintf("e.id = $%d", len(values)))
	}

	if q.UserID > 0 {
		values = append(values, q.UserID)
		conditions = append(conditions, fmt.Sprintf("e.user_id = $%d", len(values)))
	}

	if q.Status != "" {
		values = append(values, q.Status)
		conditions = append(conditions, fmt.Sprintf("e.status = $%d", len(values)))
	}

	if len(conditions) == 0 {
		return "", values
	}

	return " where " + strings.Join(conditions, " and "), values
}
