package factories

import (
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/bmviniciuss/tcc/card-payment/src/core/payment"
)

func NewPaymentService(db *pgxpool.Pool, cardApi payment.CardAPI, paymentRepository payment.PaymentRepository) payment.Service {
	return payment.NewPaymentService(cardApi, paymentRepository)
}
