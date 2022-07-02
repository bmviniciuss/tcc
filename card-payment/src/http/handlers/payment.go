package paymenthandler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/bmviniciuss/tcc/card-payment/src/core/payment"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
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
	r.Get("/", handleGetPaymentsByClientId(p.PaymentService))
	r.Post("/", handleProcessPayment(p.PaymentService))
}

type CardPaymentRequest struct {
	ClientId    string      `json:"client_id" validate:"required,uuid4"`
	Amount      float64     `json:"amount" validate:"required"`
	PaymentType string      `json:"payment_type" validate:"required"`
	PaymentDate time.Time   `json:"payment_date" validate:"required"`
	PaymentInfo PaymentInfo `json:"payment_info" validate:"required"`
}

type PaymentInfo struct {
	CardToken string `json:"card_token" validate:"required"`
}

type CardPaymentResponse struct {
	Id          string              `json:"id"`
	ClientId    string              `json:"client_id"`
	Amount      float64             `json:"amount"`
	PaymentType string              `json:"payment_type"`
	PaymentDate time.Time           `json:"payment_date"`
	PaymentInfo PaymentInfoResponse `json:"payment_info"`
}

type PaymentInfoResponse struct {
	MaskedNumber string `json:"masked_number"`
}

func handleProcessPayment(paymentService payment.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("ProcessPayment: Process started")

		w.Header().Set("Content-Type", "application/json")

		var cardPaymentRequest CardPaymentRequest

		if err := json.NewDecoder(r.Body).Decode(&cardPaymentRequest); err != nil {
			log.Println("ProcessPayment: Error deconing body validation", err)

			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		validate := validator.New()

		if err := validate.Struct(cardPaymentRequest); err != nil {
			log.Println("ProcessPayment: Error in validation", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error":   true,
				"message": err.Error(),
			})
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
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error":   true,
				"message": err.Error(),
			})
			return
		}

		cardPaymentResponse := CardPaymentResponse{
			Id:          payment.Id,
			ClientId:    payment.ClientId,
			Amount:      payment.Amount,
			PaymentType: payment.PaymentType,
			PaymentDate: payment.PaymentDate,
			PaymentInfo: PaymentInfoResponse{
				MaskedNumber: payment.PaymentInfo.MaskedNumber,
			},
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(cardPaymentResponse)
	}
}

type GetPaymentsByClientIdReponse struct {
	Content []CardPaymentResponse `json:"content"`
}

func handleGetPaymentsByClientId(paymentService payment.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("GetPaymentsByClientId: Process started")
		w.Header().Set("Content-Type", "application/json")

		clientId := r.URL.Query().Get("client_id")

		if clientId == "" {
			log.Println("GetPaymentsByClientId: no client_id was provided")
			http.Error(w, "You must provide 'client_id' in the request's url", http.StatusBadRequest)
			return
		}

		filters := &payment.GetPaymentsByClientIdInput{
			ClientId: clientId,
		}

		payments, err := paymentService.GetPaymentsByClientId(filters)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": err.Error(),
			})
			return
		}

		presentationPayments := []CardPaymentResponse{}

		for _, payment := range payments {
			p := CardPaymentResponse{
				Id:          payment.Id,
				ClientId:    payment.ClientId,
				Amount:      payment.Amount,
				PaymentType: payment.PaymentType,
				PaymentDate: payment.PaymentDate,
				PaymentInfo: PaymentInfoResponse{
					MaskedNumber: payment.PaymentInfo.MaskedNumber,
				},
			}
			presentationPayments = append(presentationPayments, p)
		}

		res := GetPaymentsByClientIdReponse{
			Content: presentationPayments,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}
