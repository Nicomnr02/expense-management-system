package expensedto

import (
	"expense-management-system/model"
	"time"
)

type FetchExpenseReq struct {
	ID     string `param:"id"`
	Status string `json:"status"`
	UserID int    `json:"user_id"`
	model.Pagination
	Timestamp time.Time
}

type FetchExpenseRes struct {
	ID          string    `json:"id"`
	UserID      int       `json:"user_id"`
	UserName    string    `json:"user_name"`
	AmountIDR   string    `json:"amount_idr"`
	Status      string    `json:"status"`
	SubmittedAt time.Time `json:"submitted_at"`
}
