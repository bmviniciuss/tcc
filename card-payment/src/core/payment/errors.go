package payment

import "errors"

var (
	ErrInvalidPaymentType        = errors.New("invalid payment type")
	ErrFetchingCard              = errors.New("error while fetching card information")
	ErrNoCardFound               = errors.New("the provided card was not found")
	ErrInvalidCardTypeForPayment = errors.New("the provided card type is not valid for this payment type")
	ErrPaymentAuthorization      = errors.New("error while processing payment authorization")
	ErrPaymentNotAuthorized      = errors.New("payment was not authorized")
)
