package expensedomain

import (
	"time"

	"github.com/google/uuid"
)

type Approval struct {
	ID           int
	ExpenseID    uuid.UUID
	ApproverID   int
	ApproverName string
	ApproverRole string
	Status       string
	Notes        string
	CreatedAt    time.Time
}
