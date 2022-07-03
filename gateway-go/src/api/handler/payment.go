package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	http_utils "github.com/bmviniciuss/gateway/src/api/utils"
	"github.com/bmviniciuss/gateway/src/core/card_payment"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func MakePaymentHandlers(r chi.Router, cardPaymentService card_payment.Service) {
	r.Post("/card", createCardPayment(cardPaymentService))
	r.Get("/card", getCardPaymentsByClientId(cardPaymentService))
}

type CreateCardPaymentRequest struct {
	ClientId    string             `json:"client_id" validate:"required,uuid4"`
	PaymentType string             `json:"payment_type" validate:"required"`
	PaymentDate string             `json:"payment_date" validate:"required"`
	Amount      float64            `json:"amount" validate:"required"`
	PaymentInfo PaymentInfoRequest `json:"payment_info" validate:"required"`
}

type PaymentInfoRequest struct {
	CardToken string `json:"card_token" validate:"required"`
}

func (c *CreateCardPaymentRequest) Validate() []error {
	validate := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(c)
	return translateError(err, trans)
}

func translateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := errors.New(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}

var (
	CreateCardPaymentServerInternalError = errors.New("Internal server error while creating card")
	GetPaymentsServerInternalError       = errors.New("Internal server error while fecthing card payments")
)

func createCardPayment(cardPaymentService card_payment.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := &CreateCardPaymentRequest{}

		err := http_utils.DecodeRequestBodyToJson(r.Body, input)

		if err != nil {
			// log.Println("Error while decoding create card payment request")
			// log.Println(err)
			http_utils.SetErrorResponse(w, http.StatusInternalServerError, CreateCardPaymentServerInternalError)
			return
		}

		if errs := input.Validate(); len(errs) > 0 {
			// log.Println("Validation error for request body")
			// log.Println(errs)
			http_utils.SetErrorResponse(w, http.StatusBadRequest, errs[0])
			return
		}

		payment := &card_payment.CardPayment{
			ClientId:    input.ClientId,
			Amount:      input.Amount,
			PaymentType: input.PaymentType,
			PaymentDate: input.PaymentDate,
			PaymentInfo: card_payment.CardPaymentInfo{
				CardToken: input.PaymentInfo.CardToken,
			},
		}

		err = cardPaymentService.CreatePayment(payment)
		if err != nil {
			// log.Println("Error while calling create card payment service")
			// log.Println("Error: ", err)
			http_utils.SetErrorResponse(w, http.StatusInternalServerError, CreateCardPaymentServerInternalError)
			return
		}

		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(payment); err != nil {
			// log.Println("Error while enconding create card response")
			// log.Println(err)
			http_utils.SetErrorResponse(w, http.StatusInternalServerError, CreateCardServerInternalError)
		}
	}
}

func getCardPaymentsByClientId(cardPaymentService card_payment.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("GetPaymentsByClientId: Process started")
		w.Header().Set("Content-Type", "application/json")

		clientId := r.URL.Query().Get("client_id")

		if clientId == "" {
			log.Println("GetPaymentsByClientId: no client_id was provided")
			http.Error(w, "You must provide 'client_id' in the request's url", http.StatusBadRequest)
			return
		}

		res, err := cardPaymentService.GetPaymentsByClientId(clientId)

		if err != nil {
			log.Println("Error while calling card payment services to get payments", err)
			http_utils.SetErrorResponse(w, http.StatusInternalServerError, GetPaymentsServerInternalError)
			return
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Println("Error while enconding get card payments response")
			log.Println(err)
			http_utils.SetErrorResponse(w, http.StatusInternalServerError, GetPaymentsServerInternalError)
		}
	}
}
