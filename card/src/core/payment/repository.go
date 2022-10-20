package payment

type PaymentAuthorizationRepository interface {
	Create(PaymentAuthorization *PaymentAuthorization) error
}
