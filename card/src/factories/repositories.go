package factories

import (
	"log"
	"os"

	postgrescardrepository "github.com/bmviniciuss/tcc/card/src/adapter/card"
	memorycardrepository "github.com/bmviniciuss/tcc/card/src/adapter/card/memory"
	postgrespaymentrepository "github.com/bmviniciuss/tcc/card/src/adapter/payment"
	memorypaymentrepository "github.com/bmviniciuss/tcc/card/src/adapter/payment/memory"

	"github.com/bmviniciuss/tcc/card/src/core/card"
	"github.com/bmviniciuss/tcc/card/src/core/payment"
	"github.com/jackc/pgx/v4/pgxpool"
)

func CardRepositoryFactory(db *pgxpool.Pool) card.CardRepository {
	useDb := os.Getenv("USE_DB")

	if useDb == "true" {
		log.Println("Use PostgresCardRepository")
		return postgrescardrepository.NewPostgresCardRepository(db)
	}
	log.Println("Use MemoryCardRepository")

	return memorycardrepository.NewMemoryCardRepository()
}

func PaymentAuthorizationRepositoryFactory(db *pgxpool.Pool) payment.PaymentAuthorizationRepository {
	useDb := os.Getenv("USE_DB")

	if useDb == "true" {
		log.Println("Use PostgresPaymentRepository")
		return postgrespaymentrepository.NewPostgresPaymentRepository(db)
	}
	log.Println("Use MemoryPaymentRepository")
	return memorypaymentrepository.NewMemoryPaymentRepository()
}
