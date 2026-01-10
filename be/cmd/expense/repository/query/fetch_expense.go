package expensequery

import "expense-management-system/model"

type FetchExpense struct {
	ID     string
	Status string
	UserID int
	model.Pagination
}
