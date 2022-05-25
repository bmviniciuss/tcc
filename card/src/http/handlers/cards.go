package handlers

import (
	"encoding/json"
	"errors"
	"io"
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
	IsCredit       *bool  `json:"is_credit" validate:"required"`
	IsDebit        *bool  `json:"is_debit" valid:"required"`
}

func handleCreateCard(cardService *card.CardService) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Calling POST /cards")

		rw.Header().Set("Content-Type", "application/json")
		validate := validator.New()
		var createCardRequest CreateCardRequest

		log.Println("[handleCreateCard] Decoding request body")
		if err := json.NewDecoder(r.Body).Decode(&createCardRequest); err != nil {
			log.Println("[handleCreateCard] Error decoding request body:", err)
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

		log.Println("[handleCreateCard] Validating request body")
		err := validate.Struct(createCardRequest)

		if err != nil {
			log.Println("[handleCreateCard] Error validating request body:", err)
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(map[string]string{"error": err.Error()})
			return
		}

		log.Println("[handleCreateCard] Generating card")
		card, err := cardService.Generate(&card.GenerateCardServiceInput{
			CardholderName: createCardRequest.CardholderName,
			IsCredit:       *createCardRequest.IsCredit,
			IsDebit:        *createCardRequest.IsDebit,
		})

		if err != nil {
			log.Println("[handleCreateCard] Error generating card:", err)
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(map[string]string{"error": err.Error()})
			return
		}

		log.Println("[handleCreateCard] Card generated")
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
		log.Println("[handleGetCard] Getting card with token:", cardToken)

		if cardToken == "" {
			log.Println("[handleGetCard] Error getting card: token is empty")
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(map[string]string{"error": "Token is required"})
			return
		}

		log.Println("[handleGetCard] Getting card")
		card, err := cardService.GetByToken(cardToken)

		if err != nil {
			log.Println("[handleGetCard] Error getting card:", err)
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(map[string]string{"error": err.Error()})
			return
		}

		log.Println("[handleGetCard] Card found")
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
