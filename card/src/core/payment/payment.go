package payment

import (
	"errors"
	"log"
	"time"

	"github.com/bmviniciuss/tcc/card/src/core/card"
)

type PaymentAuthorization struct {
	Id              string    `json:"id"`
	Amount          float64   `json:"amount"`
	Status          string    `json:"status"`
	TransactionDate time.Time `json:"transaction_date"`
	CardId          string    `json:"-"`
	CreatedAt       time.Time `json:"create_at"`
}

type CreatePaymentAuthorization struct {
	Amount          float64
	CardToken       string
	PaymentType     string
	TransactionDate time.Time
}

type PaymentService struct {
	cardRepository                 card.CardRepository
	paymentAuthorizationRepository PaymentAuthorizationRepository
}

func NewPaymentService(cadRepository card.CardRepository, paymentAuthorizationRepository PaymentAuthorizationRepository) *PaymentService {
	return &PaymentService{
		cardRepository:                 cadRepository,
		paymentAuthorizationRepository: paymentAuthorizationRepository,
	}
}

func (p *PaymentService) Authorize(input *CreatePaymentAuthorization) (*PaymentAuthorization, error) {
	log.Println("[PaymentService] Auhtorizing Payment")

	if input.PaymentType != "CREDIT" && input.PaymentType != "DEBIT" {
		return &PaymentAuthorization{}, errors.New("Invalid Payment Type")
	}

	card, err := p.cardRepository.GetByToken(input.CardToken)

	if err != err {
		log.Println("Error = ", err.Error())
		return &PaymentAuthorization{}, errors.New("Error while processing payment")
	}

	if card.Id == "" {
		return &PaymentAuthorization{}, errors.New("Card Not Found")
	}

	log.Println("Card found")

	paymentAuthorization := &PaymentAuthorization{
		Amount:          input.Amount,
		Status:          getAuthorizationStatus(input, card),
		TransactionDate: input.TransactionDate,
		CardId:          card.Id,
	}

	err = p.paymentAuthorizationRepository.Create(paymentAuthorization)

	if err != nil {
		log.Println("Error = ", err.Error())
		return &PaymentAuthorization{}, errors.New("Unable to auhtorize payment")
	}

	return paymentAuthorization, nil
}

func getAuthorizationStatus(input *CreatePaymentAuthorization, card *card.Card) string {
	if card.Active == false {
		return "DECLINED"
	}

	if input.PaymentType == "CREDIT" {
		if card.IsCredit == false {
			return "DECLINED"
		}
	} else if input.PaymentType == "DEBIT" {
		if card.IsDebit == false {
			return "DECLINED"
		}
	}

	return "APPROVED"
}
