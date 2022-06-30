package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	http_utils "github.com/bmviniciuss/gateway/src/api/utils"
	"github.com/go-chi/chi/v5"
)

func MakePaymentHandlers(r chi.Router) {
	r.Post("/card", createCardPayment())
}

type CreateCardPaymentRequest struct {
	ClientId    string             `json:"client_id"`
	PaymentType string             `json:"payment_type"`
	PaymentDate string             `json:"payment_date"`
	Amount      float64            `json:"amount"`
	PaymentInfo PaymentInfoRequest `json:"payment_info"`
}

type PaymentInfoRequest struct {
	CardToken string `json:"card_token"`
}

var (
	CreateCardPaymentServerInternalError = errors.New("Internal server error while creating card")
)

func createCardPayment() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := &CreateCardPaymentRequest{}

		err := http_utils.DecodeRequestBodyToJson(r.Body, input)

		if err != nil {
			log.Println("Error while decoding create card payment request")
			log.Println(err)
			http_utils.SetErrorResponse(w, http.StatusInternalServerError, CreateCardPaymentServerInternalError)
			return
		}

		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(input); err != nil {
			log.Println("Error while enconding create card response")
			log.Println(err)
			http_utils.SetErrorResponse(w, http.StatusInternalServerError, CreateCardServerInternalError)
		}
	}
}
