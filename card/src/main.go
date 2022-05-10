package main

import (
	"log"
	"os"

	postgrescardrepository "github.com/bmviniciuss/tcc/card/src/adapter/card"
	carddetailsgenerator "github.com/bmviniciuss/tcc/card/src/adapter/carddetails"
	"github.com/bmviniciuss/tcc/card/src/core/card"
	carddetails "github.com/bmviniciuss/tcc/card/src/core/cardDetails"
	"github.com/bmviniciuss/tcc/card/src/core/encrypter"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := "host=localhost user=root password=root dbname=card-ms port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database")
	}

	error := db.AutoMigrate(&postgrescardrepository.PostgresCard{})
	if error != nil {
		log.Fatal("Error auto migrating", err)
	}

	cardRepository := postgrescardrepository.NewPostgresCardRepository(db)

	encrypter := encrypter.NewEncrypter([]byte(os.Getenv("ENCRYPTION_KEY")))
	cardDetailsGenerator := carddetailsgenerator.NewCardDetailsGenerator()
	cardDetailsGeneratorService := carddetails.NewCardDetailsGeneratorService(cardDetailsGenerator)
	cardService := card.NewCardService(cardDetailsGeneratorService, encrypter, cardRepository)

	cardService.Generate(&card.GenerateCardServiceInput{
		CardholderName: "Vinicius",
		IsCredit:       true,
		IsDebit:        true,
	})
}
