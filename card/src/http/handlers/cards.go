package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bmviniciuss/tcc/card/src/core/card"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type CardsController struct {
	CardService *card.CardService
}

func NewCardsController(cardService *card.CardService) CardsController {
	return CardsController{
		CardService: cardService,
	}
}

func (c CardsController) Route(r chi.Router) {
	r.Post("/", handleCreateCard(c.CardService))
	r.Get("/", handleGetCard(c.CardService))
}

type PresentationCard struct {
	Id              string `json:"id"`
	CardholderName  string `json:"cardholder_name"`
	Token           string `json:"token"`
	MaskedNumber    string `json:"masked_number"`
	ExpirationYear  int    `json:"expiration_year"`
	ExpirationMonth int    `json:"expiration_month"`
	Active          bool   `json:"active"`
	IsCredit        bool   `json:"is_credit"`
	IsDebit         bool   `json:"is_debit"`
}

type CreateCardRequest struct {
	CardholderName string `json:"cardholder_name" validate:"required"`
	IsCredit       bool   `json:"is_credit" validate:"required"`
	IsDebit        bool   `json:"is_debit" valid:"required"`
}

func handleCreateCard(cardService *card.CardService) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		validate := validator.New()
		rw.Header().Set("Content-Type", "application/json")
		log.Println("Calling POST /cards")
		var createCardRequest CreateCardRequest

		if err := json.NewDecoder(r.Body).Decode(&createCardRequest); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(map[string]string{"error": err.Error()})
			return
		}

		err := validate.Struct(createCardRequest)

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(map[string]string{"error": err.Error()})
			return
		}

		card, err := cardService.Generate(&card.GenerateCardServiceInput{
			CardholderName: createCardRequest.CardholderName,
			IsCredit:       createCardRequest.IsCredit,
			IsDebit:        createCardRequest.IsDebit,
		})

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(map[string]string{"error": err.Error()})
			return
		}

		presentationCard := parseCardToPresentationCard(card)

		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(presentationCard)
	}
}

func handleGetCard(cardService *card.CardService) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		log.Println("Calling GET /cards")
		cardToken := r.URL.Query().Get("token")
		log.Println("Card token:", cardToken)

		if cardToken == "" {
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(map[string]string{"error": "Token is required"})
			return
		}

		card, err := cardService.GetByToken(cardToken)

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(map[string]string{"error": err.Error()})
			return
		}

		presentationCard := parseCardToPresentationCard(card)

		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(map[string]interface{}{
			"data": presentationCard,
		})
	}
}

func parseCardToPresentationCard(card *card.Card) PresentationCard {
	return PresentationCard{
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
}
