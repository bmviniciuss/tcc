package handlers

type PaymentAuthorization struct {
	Id              string  `json:"id"`
	Amount          float64 `json:"amount"`
	State           string  `json:"status"`
	TransactionDate string  `json:"transaction_date"`
	CreatedAt       string  `json:"created_at"`
}

type CreatePaymentAuthorization struct {
	ServiceId       string  `json:"service_id"`
	Amount          float64 `json:"amount"`
	CardToken       string  `json:"card_token"`
	PaymentType     string  `json:"payment_type"`
	TransactionDate string  `json:"transaction_date"`
}
