package memorypaymentrepository

import (
	"time"

	"github.com/bmviniciuss/tcc/card/src/core/payment"
	"github.com/google/uuid"
)

type memoryPaymentRepository struct {
}

func NewPostgresPaymentRepository() *memoryPaymentRepository {
	return &memoryPaymentRepository{}
}

func (r *memoryPaymentRepository) Create(paymentAuthorization *payment.PaymentAuthorization) error {
	paymentAuthorization.Id = uuid.NewString()
	paymentAuthorization.CreatedAt = time.Now()

	return nil
}
