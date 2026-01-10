package expensedto

import "time"

type ApprovalReq struct {
	ID        string
	Action    string
	Notes     string `json:"notes" validate:"required,min=5,max=500"`
	Timestamp time.Time
}

type ApprovalRes struct {
	Status string `json:"status"`
}
