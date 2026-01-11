package expenserepository

import (
	"context"
	expensedomain "expense-management-system/cmd/expense/domain"
	"expense-management-system/database"
)

func (r *expenseRepositoryImpl) CreateApproval(c context.Context, tx database.Tx, data expensedomain.Approval) error {
	query := `
	insert into approvals (
        expense_id,
        approver_id,
        status,
        notes,
        created_at
    ) values ($1, $2, $3, $4, $5)
	`

	_, err := tx.Exec(c, query,
		data.ExpenseID,
		data.ApproverID,
		data.Status,
		data.Notes,
		data.CreatedAt,
	)
	
	return err
}
