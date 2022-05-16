package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bmviniciuss/tcc/card/src/adapter/db"
	api "github.com/bmviniciuss/tcc/card/src/http"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// dsn := "host=localhost user=root password=root dbname=cards-ms port=5432"
	db := db.ConnectDB()

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
