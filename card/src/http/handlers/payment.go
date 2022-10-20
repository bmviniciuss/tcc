package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/bmviniciuss/tcc/card/src/core/payment"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type PaymentController struct {
	PaymentService *payment.PaymentService
}

func NewPaymentController(paymentService *payment.PaymentService) PaymentController {
	return PaymentController{
		PaymentService: paymentService,
	}

}

func (c PaymentController) Route(r chi.Router) {
	r.Post("/authorize", handleAuthorizePayment(c.PaymentService))
}

func handleAuthorizePayment(paymentService *payment.PaymentService) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("POST /authorize")

		rw.Header().Set("Content-Type", "application/json")
		validate := validator.New()
		var createPaymentAuthorization CreatePaymentAuthorization

		if err := json.NewDecoder(r.Body).Decode(&createPaymentAuthorization); err != nil {
			log.Println("[handleAuthorizePayment] Error decoding request body:", err)
			rw.WriteHeader(http.StatusBadRequest)

			errorMessage := ""
			if errors.Is(err, io.EOF) {
				errorMessage = "Request body is empty"
			} else {
				errorMessage = err.Error()
			}

			json.NewEncoder(rw).Encode(map[string]string{"error": errorMessage})
			return
		}

		log.Println("[handleAuthorizePayment] Validating request body")
		err := validate.Struct(createPaymentAuthorization)

		if err != nil {
			log.Println("[handleAuthorizePayment] Error validating request body:", err)
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(map[string]string{"error": err.Error()})
			return
		}

		time, err := time.Parse(time.RFC3339, createPaymentAuthorization.TransactionDate)
		if err != nil {
			log.Println("[handleAuthorizePayment] Error Parsing date request body:", err)
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(map[string]string{"error": err.Error()})
			return
		}

		log.Println("[handleAuthorizePayment] Generating card")

		paymentAuthorization, err := paymentService.Authorize(&payment.CreatePaymentAuthorization{
			Amount:          createPaymentAuthorization.Amount,
			CardToken:       createPaymentAuthorization.CardToken,
			PaymentType:     createPaymentAuthorization.PaymentType,
			TransactionDate: time,
		})

		if err != nil {
			log.Println("[handleAuthorizePayment] Error processing payment:", err)
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(map[string]string{"error": err.Error()})
			return
		}

		p := &PaymentAuthorizationPresentation{
			Id:              paymentAuthorization.Id,
			Amount:          paymentAuthorization.Amount,
			Status:          paymentAuthorization.Status,
			TransactionDate: paymentAuthorization.TransactionDate.Format("2006-01-02T15:04:05Z07:00"),
			CreateAt:        paymentAuthorization.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}

		rw.WriteHeader(http.StatusCreated)
		json.NewEncoder(rw).Encode(p)
	}
}

type PaymentAuthorizationPresentation struct {
	Id              string  `json:"id"`
	Amount          float64 `json:"amount"`
	Status          string  `json:"status"`
	TransactionDate string  `json:"transaction_date"`
	CreateAt        string  `json:"create_at"`
}
