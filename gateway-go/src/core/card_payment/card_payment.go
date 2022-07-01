package card_payment

type CardPayment struct {
	Id          string          `json:"id"`
	ClientId    string          `json:"client_id"`
	Amount      float64         `json:"amount"`
	PaymentType string          `json:"payment_type"`
	PaymentDate string          `json:"payment_date"`
	PaymentInfo CardPaymentInfo `json:"payment_info"`
}

type CardPaymentInfo struct {
	MaskedNumber string `json:"masked_number"`
	CardToken    string `json:"-"`
}

type Service interface {
	CreatePayment(payment *CardPayment) error
}
