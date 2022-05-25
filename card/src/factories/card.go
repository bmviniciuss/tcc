package factories

import (
	"os"

	postgrescardrepository "github.com/bmviniciuss/tcc/card/src/adapter/card"
	carddetailsgenerator "github.com/bmviniciuss/tcc/card/src/adapter/carddetails"
	"github.com/bmviniciuss/tcc/card/src/core/card"
	"github.com/bmviniciuss/tcc/card/src/core/encrypter"
	"github.com/jmoiron/sqlx"
)

func CardServiceFactory(db *sqlx.DB) *card.CardService {
	cardRepository := postgrescardrepository.NewPostgresCardRepository(db)
	encrypter := encrypter.NewEncrypter([]byte(os.Getenv("ENCRYPTION_KEY")))
	cardDetailsGenerator := carddetailsgenerator.NewCardDetailsGenerator()
	cardService := card.NewCardService(cardDetailsGenerator, encrypter, cardRepository)
	return cardService
}
