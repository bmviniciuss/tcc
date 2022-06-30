package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/bmviniciuss/gateway/src/api/presenter"
	http_utils "github.com/bmviniciuss/gateway/src/api/utils"
	"github.com/bmviniciuss/gateway/src/core/card"
	"github.com/go-chi/chi/v5"
)

func MakeCardHandlers(r chi.Router, cardService card.Service) {
	r.Post("/", createCard(cardService))
}

var (
	CreateCardServerInternalError = errors.New("Internal server error while creating card")
)

type CreateCardInput struct {
	CardHolderName string `json:"cardholder_name"`
	IsCredit       bool   `json:"is_credit"`
	IsDebit        bool   `json:"is_debit"`
}

func createCard(cardService card.Service) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		input := &CreateCardInput{}

		err := http_utils.DecodeRequestBodyToJson(r.Body, input)

		if err != nil {
			log.Println("Error while decoding create card request")
			log.Println(err)
			http_utils.SetErrorResponse(w, http.StatusInternalServerError, CreateCardServerInternalError)
			return
		}

		// TODO: add struct validation

		card, err := cardService.CreateCard(&card.CreateCardRequest{
			CardholderName: input.CardHolderName,
			IsCredit:       &input.IsCredit,
			IsDebit:        &input.IsDebit,
		})

		if err != nil {
			log.Println("Error while calling create card service")
			log.Println(err)
			http_utils.SetErrorResponse(w, http.StatusInternalServerError, CreateCardServerInternalError)
			return
		}

		p := presenter.Card{
			Id:              card.Id,
			CardholderName:  card.CardholderName,
			Token:           card.Token,
			MaskedNumber:    card.MaskedNumber,
			ExpirationYear:  card.ExpirationYear,
			ExpirationMonth: card.ExpirationMonth,
			Active:          card.Active,
			IsCredit:        card.IsCredit,
			IsDebit:         card.IsDebit,
		}
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(p); err != nil {
			log.Println("Error while enconding create card response")
			log.Println(err)
			http_utils.SetErrorResponse(w, http.StatusInternalServerError, CreateCardServerInternalError)
		}
	})
}
