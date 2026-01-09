package expensedomain

import (
	"time"

	"github.com/google/uuid"
)

type Expense struct {
	ID          uuid.UUID
	UserID      int
	UserName    string
	Amount      int
	Description string
	ReceiptURL  string
	Status      string
	SubmittedAt time.Time
	ProcessedAt time.Time
}
