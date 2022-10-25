package factories

import (
	postgrescardrepository "github.com/bmviniciuss/tcc/card/src/adapter/card/memory"
	carddetailsgenerator "github.com/bmviniciuss/tcc/card/src/adapter/carddetails"
	postgrespaymentrepository "github.com/bmviniciuss/tcc/card/src/adapter/payment/memory"
	"github.com/bmviniciuss/tcc/card/src/core/card"
	"github.com/bmviniciuss/tcc/card/src/core/payment"
	"github.com/jackc/pgx/v4/pgxpool"
)

func CardServiceFactory(db *pgxpool.Pool) *card.CardService {
	cardRepository := postgrescardrepository.NewPostgresCardRepository()
	cardDetailsGenerator := carddetailsgenerator.NewCardDetailsGenerator()
	cardService := card.NewCardService(cardDetailsGenerator, cardRepository)
	return cardService
}

func PaymentServiceFactory(db *pgxpool.Pool) *payment.PaymentService {
	cardRepository := postgrescardrepository.NewPostgresCardRepository()
	paymentRepository := postgrespaymentrepository.NewPostgresPaymentRepository()
	return payment.NewPaymentService(cardRepository, paymentRepository)
}
