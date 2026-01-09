package expensequery

type FetchExpense struct {
	ID     string
	Status string
	UserID int
	Limit  int
	Offset int
}
