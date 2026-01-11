package expensedomain

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID         int
	ExternalID uuid.UUID
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
