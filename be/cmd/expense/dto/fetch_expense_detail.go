package expensedto

import (
	"time"

	"github.com/google/uuid"
)

type FetchExpenseDetailReq struct {
	ID        string `json:"id"`
	Timestamp time.Time
}

type FetchExpenseDetailRes struct {
	ID          uuid.UUID `json:"id"`
	UserID      int       `json:"user_id"`
	UserName    string    `json:"user_name"`
	AmountIDR   string    `json:"amount_idr"`
	Description string    `json:"description"`
	ReceiptURL  string    `json:"receipt_url"`
	Status      string    `json:"status"`
	SubmittedAt time.Time `json:"submitted_at"`
	ProcessedAt time.Time `json:"processed_at"`
}
