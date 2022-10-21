package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bmviniciuss/gateway/src/api/router"
)

func main() {
	// err := godotenv.Load()

	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	grpcEnabled := os.Getenv("GRPC_ENABLED") == "true"

	log.Println("Starting up Gateway Server...")

	if grpcEnabled == true {
		log.Println("Mode: gRPC")
	} else {
		log.Println("Mode: HTTP")
	}

	r := router.GetRouter()

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Println("failed to start server", err)
		os.Exit(1)
	}

	log.Println("ready to serve requests on " + addr)
}
