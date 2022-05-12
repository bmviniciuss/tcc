package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	postgrescardrepository "github.com/bmviniciuss/tcc/card/src/adapter/card"
	carddetailsgenerator "github.com/bmviniciuss/tcc/card/src/adapter/carddetails"
	"github.com/bmviniciuss/tcc/card/src/core/card"
	"github.com/bmviniciuss/tcc/card/src/core/encrypter"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CardsController struct {
	Db *gorm.DB
}

func NewCardsController(db *gorm.DB) CardsController {
	return CardsController{
		Db: db,
	}
}

func (c CardsController) Route(r chi.Router) {
	r.Post("/", handleCreateCard(c.Db))
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
	CardholderName string `json:"cardholder_name" validate:"required"`
	IsCredit       *bool  `json:"is_credit" validate:"required"`
	IsDebit        *bool  `json:"is_debit" valid:"required"`
}

func handleCreateCard(db *gorm.DB) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		validate := validator.New()

		rw.Header().Set("Content-Type", "application/json")

		log.Println("Calling POST /cards")
		var createCardRequest CreateCardRequest
		fmt.Println(r.Body)

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

		cardRepository := postgrescardrepository.NewPostgresCardRepository(db)
		encrypter := encrypter.NewEncrypter([]byte("gFvJR96@UXYrq_2m"))
		cardDetailsGenerator := carddetailsgenerator.NewCardDetailsGenerator()
		cardService := card.NewCardService(cardDetailsGenerator, encrypter, cardRepository)

		card, err := cardService.Generate(&card.GenerateCardServiceInput{
			CardholderName: createCardRequest.CardholderName,
			IsCredit:       *createCardRequest.IsCredit,
			IsDebit:        *createCardRequest.IsDebit,
		})

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(rw).Encode(map[string]string{"error": err.Error()})
			return
		}

		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(card)
	}
}
