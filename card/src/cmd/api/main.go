package main

import (
	"log"
	"net/http"
	"os"

	api "github.com/bmviniciuss/tcc/card/src/http"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appPort := os.Getenv("PORT")
	mux := api.NewMux()

	server := http.Server{
		Addr:    ":" + appPort,
		Handler: mux,
	}

	log.Println("Server started on port " + appPort)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("Server closed unexpected")
	}

}
