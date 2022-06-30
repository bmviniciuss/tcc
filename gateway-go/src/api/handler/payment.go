package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	http_utils "github.com/bmviniciuss/gateway/src/api/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func MakePaymentHandlers(r chi.Router) {
	r.Post("/card", createCardPayment())
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
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
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

		if errs := input.Validate(); len(errs) > 0 {
			log.Println("Validation error for request body")
			log.Println(err)
			http_utils.SetErrorResponse(w, http.StatusInternalServerError, errs[0])
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
