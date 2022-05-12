package main

import (
	"log"
	"net/http"
	"os"

	postgrescardrepository "github.com/bmviniciuss/tcc/card/src/adapter/card"
	api "github.com/bmviniciuss/tcc/card/src/http"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := "host=localhost user=root password=root dbname=cards-ms port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&postgrescardrepository.PostgresCard{})

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")

	appPort := os.Getenv("PORT")
	mux := api.NewApi(db)

	server := http.Server{
		Addr:    ":" + appPort,
		Handler: mux,
	}

	log.Println("Server started on port " + appPort)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("Server closed unexpected")
	}

}
