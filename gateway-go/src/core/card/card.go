package card

type PresentationCard struct {
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

type CreateCardRequest struct {
	CardholderName string `json:"cardholder_name" validate:"required"`
	IsCredit       *bool  `json:"is_credit" validate:"required"`
	IsDebit        *bool  `json:"is_debit" valid:"required"`
}

type Service interface {
	CreateCard(input *CreateCardRequest) (*PresentationCard, error)
}
