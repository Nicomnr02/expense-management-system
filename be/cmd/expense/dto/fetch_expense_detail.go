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
	Expense   FetchExpenseDetailExpenseRes   `json:"expense"`
	Appproval *FetchExpenseDetailApprovalRes `json:"approval"`
	Payment   *FetchExpenseDetailPaymentRes  `json:"payment"`
}

type FetchExpenseDetailExpenseRes struct {
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

type FetchExpenseDetailApprovalRes struct {
	ID           int       `json:"id"`
	ApproverID   int       `json:"approver_id"`
	ApproverName string    `json:"approver_name"`
	ApproverRole string    `json:"approver_role"`
	Status       string    `json:"status"`
	Notes        string    `json:"notes"`
	CreatedAt    time.Time `json:"created_at"`
}

type FetchExpenseDetailPaymentRes struct {
	ID         int       `json:"id"`
	ExternalID uuid.UUID `json:"external_id"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
