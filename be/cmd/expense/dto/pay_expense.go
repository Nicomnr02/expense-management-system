package expensedto

type PaymentReq struct {
	Amount     int    `json:"amount"`
	ExternalID string `json:"external_id"`
}

type PaymentRes struct {
	Data    PaymentDataRes `json:"data"`
	Message string         `json:"message"`
}

type PaymentDataRes struct {
	ID         string `json:"id"`
	ExternalID string `json:"external_id"`
	Status     string `json:"status"`
}
