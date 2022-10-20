package factories

import (
	postgrescardrepository "github.com/bmviniciuss/tcc/card/src/adapter/card"
	carddetailsgenerator "github.com/bmviniciuss/tcc/card/src/adapter/carddetails"
	postgrespaymentrepository "github.com/bmviniciuss/tcc/card/src/adapter/payment"
	"github.com/bmviniciuss/tcc/card/src/core/card"
	"github.com/bmviniciuss/tcc/card/src/core/payment"
	"github.com/jackc/pgx/v4/pgxpool"
)

func CardServiceFactory(db *pgxpool.Pool) *card.CardService {
	cardRepository := postgrescardrepository.NewPostgresCardRepository(db)
	cardDetailsGenerator := carddetailsgenerator.NewCardDetailsGenerator()
	cardService := card.NewCardService(cardDetailsGenerator, cardRepository)
	return cardService
}

func PaymentServiceFactory(db *pgxpool.Pool) *payment.PaymentService {
	cardRepository := postgrescardrepository.NewPostgresCardRepository(db)
	paymentRepository := postgrespaymentrepository.NewPostgresPaymentRepository(db)
	return payment.NewPaymentService(cardRepository, paymentRepository)
}
