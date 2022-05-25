package payment

type PaymentRepository interface {
	Create(payment *Payment) error
}
