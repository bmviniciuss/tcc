package payment

import "time"

const (
	CREDIT_CARD = "CREDIT_CARD"
	DEBIT_CARD  = "DEBIT_CARD"
)

type Payment struct {
	Id              string
	ClientId        string
	Amount          float64
	PaymentType     string
	PaymentInfo     PaymentInfo
	PaymentDate     time.Time
	Payable         Payable
	AuthorizationId string
}

type Amount struct {
	Value    float64
	Currency string
}

type PaymentInfo struct {
	CardholderName string
	CardToken      string
	MaskedNumber   string
}

type Card struct {
	Id              string `json:"id"`
	CardholderName  string `json:"cardholder_name"`
	Token           string `json:"token"`
	MaskedNumber    string `json:"masked_number"`
	ExpirationYear  int    `json:"expiration_year"`
	ExpirationMonth int    `json:"expiration_month"`
	Active          bool   `json:"active"`
	IsCredit        bool   `json:"is_credit"`
	IsDebit         bool   `json:"is_debit"`
}

type Payable struct {
	Id          string
	ClientId    string
	PaymentId   string
	PaymentDate time.Time
	Amount      float64
}

type PaymentAuthorization struct {
	Id              string  `json:"id"`
	Amount          float64 `json:"amount"`
	Status          string  `json:"status"`
	TransactionDate string  `json:"transaction_date"`
	CreateAt        string  `json:"create_at"`
}

type PaymentAuthorizationInput struct {
	Amount          float64 `json:"amount"`
	CardToken       string  `json:"card_token"`
	PaymentType     string  `json:"payment_type"`
	TransactionDate string  `json:"transaction_date"`
}
