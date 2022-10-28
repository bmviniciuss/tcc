package factories

import (
	carddetailsgenerator "github.com/bmviniciuss/tcc/card/src/adapter/carddetails"
	"github.com/bmviniciuss/tcc/card/src/core/card"
	"github.com/bmviniciuss/tcc/card/src/core/payment"
	"github.com/jackc/pgx/v4/pgxpool"
)

func CardServiceFactory(db *pgxpool.Pool) *card.CardService {
	cardRepository := CardRepositoryFactory(db)
	cardDetailsGenerator := carddetailsgenerator.NewCardDetailsGenerator()
	cardService := card.NewCardService(cardDetailsGenerator, cardRepository)
	return cardService
}

func PaymentServiceFactory(db *pgxpool.Pool) *payment.PaymentService {
	cardRepository := CardRepositoryFactory(db)
	paymentRepository := PaymentAuthorizationRepositoryFactory(db)
	return payment.NewPaymentService(cardRepository, paymentRepository)
}
