package paymenthandler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bmviniciuss/tcc/card-payment/src/core/payment"
	"github.com/go-chi/chi/v5"
)

type PaymentController struct {
	PaymentService payment.Service
}

func NewPaymentController(paymentService payment.Service) PaymentController {
	return PaymentController{
		PaymentService: paymentService,
	}
}

func (p PaymentController) Route(r chi.Router) {
	r.Post("/", handleProcessPayment(p.PaymentService))
}

type CardPaymentRequest struct {
	ClientId    string      `json:"client_id"`
	Amount      float64     `json:"amount"`
	PaymentType string      `json:"payment_type"`
	PaymentDate time.Time   `json:"payment_date"`
	PaymentInfo PaymentInfo `json:"payment_info"`
}

type PaymentInfo struct {
	CardToken string `json:"card_token"`
}

func handleProcessPayment(paymentService payment.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var cardPaymentRequest CardPaymentRequest

		if err := json.NewDecoder(r.Body).Decode(&cardPaymentRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		input := payment.ProcessPaymentInput{
			ClientId:    cardPaymentRequest.ClientId,
			Amount:      cardPaymentRequest.Amount,
			PaymentType: cardPaymentRequest.PaymentType,
			PaymentDate: cardPaymentRequest.PaymentDate,
			PaymentInfo: payment.PaymentInfoInput{
				CardToken: cardPaymentRequest.PaymentInfo.CardToken,
			},
		}

		payment, err := paymentService.Process(&input)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(payment)
	}
}
