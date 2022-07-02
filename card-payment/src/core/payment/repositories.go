package payment

type PaymentRepository interface {
	Create(payment *Payment) error
	GetPaymentsByClientId(input *GetPaymentsByClientIdInput) ([]Payment, error)
}
