package factories

import (
	"log"
	"os"

	postgrespaymentrepository "github.com/bmviniciuss/tcc/card-payment/src/adapters/payment"
	memorymentrepository "github.com/bmviniciuss/tcc/card-payment/src/adapters/payment/memory"
	"github.com/bmviniciuss/tcc/card-payment/src/core/payment"
	"github.com/jackc/pgx/v4/pgxpool"
)

func PaymentRepositoryFactory(db *pgxpool.Pool) payment.PaymentRepository {
	useDb := os.Getenv("USE_DB")

	if useDb == "true" {
		log.Println("Use PostgresPaymentRepository")
		return postgrespaymentrepository.NewPostgresPaymentRepository(db)
	}
	log.Println("Use MemoryPaymentRepository")

	return memorymentrepository.NewPaymentRepository()
}
