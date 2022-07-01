package factories

import (
	postgrescardrepository "github.com/bmviniciuss/tcc/card/src/adapter/card"
	carddetailsgenerator "github.com/bmviniciuss/tcc/card/src/adapter/carddetails"
	"github.com/bmviniciuss/tcc/card/src/core/card"
	"github.com/jmoiron/sqlx"
)

func CardServiceFactory(db *sqlx.DB) *card.CardService {
	cardRepository := postgrescardrepository.NewPostgresCardRepository(db)
	cardDetailsGenerator := carddetailsgenerator.NewCardDetailsGenerator()
	cardService := card.NewCardService(cardDetailsGenerator, cardRepository)
	return cardService
}
