package payment

import "time"

const (
	CREDIT_CARD = "CREDIT_CARD"
	DEBIT_CARD  = "DEBIT_CARD"
)

type Payment struct {
	Id          string
	ClientId    string
	PaymentType string
	Amount      Amount
	PaymentInfo PaymentInfo
	Buyer       Buyer
	Date        time.Time
}

type Amount struct {
	Value    int
	Currency string
}

type PaymentInfo struct {
	CardholderName      string
	CardToken           string
	MaskedNumber        string
	CardExpirationYear  int
	CardExpirationMonth int
}

type Buyer struct {
	Name           string
	DocumentNumber string
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
