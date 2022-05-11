package main

import (
	"log"

	"github.com/gin-gonic/gin"

	postgrescardrepository "github.com/bmviniciuss/tcc/card/src/adapter/card"
	cardshandler "github.com/bmviniciuss/tcc/card/src/http/handlers/cards"
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

	r := gin.Default()

	cr := cardshandler.NewHandler(db)
	cr.RegisterRoutes(r)

	r.Run(":3000")
}
