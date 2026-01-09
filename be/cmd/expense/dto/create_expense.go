package expensedto

import (
	"time"

	"github.com/google/uuid"
)

type CreateExpenseReq struct {
	AmountIDR   int    `json:"amount_idr" validate:"required,gt=0"`
	Description string `json:"description" validate:"required,min=5,max=500"`
	ReceiptURL  string `json:"receipt_url" validate:"required,url"`
	Timestamp   time.Time
}

type CreateExpenseRes struct {
	ID          uuid.UUID `json:"id"`
	UserID      int       `json:"user_id"`
	UserName    string    `json:"user_name"`
	AmountIDR   int       `json:"amount_idr"`
	SubmittedAt time.Time `json:"submitted_at"`
	ProcessedAt time.Time `json:"processed_at"`
	Status      string    `json:"status"`
}
