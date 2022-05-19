package payment

import (
	"fmt"
	"time"
)

type ProcessPaymentInput struct {
	ClientId    string
	Type        string
	Amount      Amount
	PaymentInfo PaymentInfoInput
	Buyer       Buyer
	Date        time.Time
}

type PaymentInfoInput struct {
	CardholderName      string
	CardToken           string
	CardExpirationYear  int
	CardExpirationMonth int
}

type Service interface {
	Process(input *ProcessPaymentInput) (*Payment, error)
}

type CardAPI interface {
	GetCardByToken(token string) (*Card, error)
}

type PaymentService struct {
	CardAPI CardAPI
}

func NewPaymentService(cardAPI CardAPI) *PaymentService {
	return &PaymentService{
		CardAPI: cardAPI,
	}
}

func (s *PaymentService) Process(input *ProcessPaymentInput) (*Payment, error) {
	if input.Type != CREDIT_CARD && input.Type != DEBIT_CARD {
		return &Payment{}, ErrInvalidPaymentType
	}

	card, err := s.CardAPI.GetCardByToken(input.PaymentInfo.CardToken)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v\n", card)

	return &Payment{}, ErrInvalidPaymentType

}
