package payment

import (
	"log"
	"time"

	"github.com/bmviniciuss/tcc/card-payment/src/constants"
)

type ProcessPaymentInput struct {
	ClientId    string
	Amount      float64
	PaymentType string
	PaymentDate time.Time
	PaymentInfo PaymentInfoInput
}

type GetPaymentsByClientIdInput struct {
	ClientId string
}
type PaymentInfoInput struct {
	CardToken string
}

type Service interface {
	Process(input *ProcessPaymentInput) (*Payment, error)
	GetPaymentsByClientId(input *GetPaymentsByClientIdInput) ([]Payment, error)
}

type CardAPI interface {
	GetCardByToken(token string) (*Card, error)
	AuthorizePayment(input *PaymentAuthorizationInput) (*PaymentAuthorization, error)
}

type PaymentService struct {
	CardAPI           CardAPI
	PaymentRepository PaymentRepository
}

func NewPaymentService(cardAPI CardAPI, paymentRepository PaymentRepository) *PaymentService {
	return &PaymentService{
		CardAPI:           cardAPI,
		PaymentRepository: paymentRepository,
	}
}

func (s *PaymentService) Process(input *ProcessPaymentInput) (*Payment, error) {
	if input.PaymentType != CREDIT_CARD && input.PaymentType != DEBIT_CARD {
		return &Payment{}, ErrInvalidPaymentType
	}

	card, err := s.CardAPI.GetCardByToken(input.PaymentInfo.CardToken)

	if err != nil {
		return nil, ErrFetchingCard
	}

	if card == nil || card.Id == "" {
		return nil, ErrNoCardFound
	}

	if (input.PaymentType == CREDIT_CARD && !card.IsCredit) || (input.PaymentType == DEBIT_CARD && !card.IsDebit) {
		return nil, ErrInvalidCardTypeForPayment
	}

	pa := &PaymentAuthorizationInput{
		Amount:          input.Amount,
		CardToken:       input.PaymentInfo.CardToken,
		PaymentType:     s.getPaymentTypeForAuthorizer(input.PaymentType),
		TransactionDate: time.Now().Format(constants.RFC3399),
	}

	paymentAuth, err := s.CardAPI.AuthorizePayment(pa)

	if err != nil {
		log.Println("Error = ", err.Error())
		return nil, ErrPaymentAuthorization
	}

	if paymentAuth.Status == "DECLINED" {
		return nil, ErrPaymentNotAuthorized
	}

	fee := getPaymentFeeByPaymentType(input.PaymentType)

	payment := &Payment{
		ClientId:        input.ClientId,
		Amount:          input.Amount,
		PaymentType:     input.PaymentType,
		PaymentDate:     input.PaymentDate,
		AuthorizationId: paymentAuth.Id,
		PaymentInfo: PaymentInfo{
			CardholderName: card.CardholderName,
			CardToken:      card.Token,
			MaskedNumber:   card.MaskedNumber,
		},
		Payable: Payable{
			ClientId:    input.ClientId,
			Amount:      input.Amount * (1 - fee),
			PaymentDate: input.PaymentDate,
		},
	}

	err = s.PaymentRepository.Create(payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func getPaymentFeeByPaymentType(paymentType string) float64 {
	if paymentType == CREDIT_CARD {
		return CREDIT_CARD_FEE
	}
	return DEBIT_CARD_FEE
}

func (s *PaymentService) GetPaymentsByClientId(input *GetPaymentsByClientIdInput) ([]Payment, error) {
	return s.PaymentRepository.GetPaymentsByClientId(input)
}

func (s *PaymentService) getPaymentTypeForAuthorizer(paymentType string) string {
	if paymentType == "CREDIT_CARD" {
		return "CREDIT"
	} else if paymentType == "DEBIT_CARD" {
		return "DEBIT"
	} else {
		return "UNSUPPORTED"
	}
}
