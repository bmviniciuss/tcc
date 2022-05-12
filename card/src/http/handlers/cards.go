package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/go-chi/chi/v5"
)

type CardsController struct {
}

func NewCardsController() CardsController {
	return CardsController{}
}

func (c CardsController) Route(r chi.Router) {
	r.Post("/", handleCreateCard())
}

type PresentationCard struct {
	Id             string `json:"id"`
	Cvv            string `json:"cvv"`
	CardholderName string `json:"cardholder_name"`
	Token          string `json:"token"`
	MaskedNumber   string `json:"masked_number"`
	Active         bool   `json:"active"`
	IsCredit       bool   `json:"is_credit"`
	IsDebit        bool   `json:"is_debit"`
}

type CreateCardRequest struct {
	CardholderName string `json:"cardholder_name" valid:"required, notnull"`
	IsCredit       *bool  `json:"is_credit"`
	IsDebit        *bool  `json:"is_debit" valid:"required, notnull, type(bool)"`
}

func handleCreateCard() func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		log.Println("Calling POST /cards")
		var createCardRequest CreateCardRequest
		fmt.Println(r.Body)

		if err := json.NewDecoder(r.Body).Decode(&createCardRequest); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(map[string]string{"error": err.Error()})
			return
		}

		_, err := govalidator.ValidateStruct(createCardRequest)

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(map[string]string{"error": err.Error()})
			return
		}

		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(createCardRequest)
	}
}
