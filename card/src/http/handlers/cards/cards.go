package cardshandler

import (
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	postgrescardrepository "github.com/bmviniciuss/tcc/card/src/adapter/card"
	carddetailsgenerator "github.com/bmviniciuss/tcc/card/src/adapter/carddetails"
	"github.com/bmviniciuss/tcc/card/src/core/card"
	carddetails "github.com/bmviniciuss/tcc/card/src/core/cardDetails"
	"github.com/bmviniciuss/tcc/card/src/core/encrypter"
)

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

type handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *handler {
	return &handler{
		db: db,
	}
}

func (h *handler) RegisterRoutes(g *gin.Engine) {
	g.POST("/cards", h.createCard)
}

func (h *handler) createCard(c *gin.Context) {
	cardRepository := postgrescardrepository.NewPostgresCardRepository(h.db)
	encrypter := encrypter.NewEncrypter([]byte(os.Getenv("ENCRYPTION_KEY")))
	cardDetailsGenerator := carddetailsgenerator.NewCardDetailsGenerator()
	cardDetailsGeneratorService := carddetails.NewCardDetailsGeneratorService(cardDetailsGenerator)
	cardService := card.NewCardService(cardDetailsGeneratorService, encrypter, cardRepository)

	card, err := cardService.Generate(&card.GenerateCardServiceInput{
		CardholderName: "Vinicius",
		IsCredit:       true,
		IsDebit:        true,
	})

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"card": fromCardToPresentation(card),
		})
	}
}

func fromCardToPresentation(card *card.Card) *PresentationCard {
	return &PresentationCard{
		Id:             card.Id,
		Cvv:            card.Cvv,
		CardholderName: card.CardholderName,
		Token:          card.Token,
		MaskedNumber:   card.MaskedNumber,
		Active:         card.Active,
		IsCredit:       card.IsCredit,
		IsDebit:        card.IsDebit,
	}
}
