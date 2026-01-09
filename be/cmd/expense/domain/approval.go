package expensedomain

import (
	"time"

	"github.com/google/uuid"
)

type Approval struct {
	ID         int
	ExpenseID  uuid.UUID
	ApproverID int
	Status     string
	Notes      string
	CreatedAt  time.Time
}
