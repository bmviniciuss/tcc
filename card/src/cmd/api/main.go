package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/bmviniciuss/tcc/card/src/adapter/db"
	"github.com/bmviniciuss/tcc/card/src/core/card"
	"github.com/bmviniciuss/tcc/card/src/core/payment"
	"github.com/bmviniciuss/tcc/card/src/factories"
	grpccard "github.com/bmviniciuss/tcc/card/src/grpc"
	"github.com/bmviniciuss/tcc/card/src/grpc/pb"
	api "github.com/bmviniciuss/tcc/card/src/http"
	"github.com/jackc/pgx/v4/pgxpool"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// err := godotenv.Load()

	// if err != nil {
	// 	log.Fatal("[Main] Error loading .env file")
	// }

	db := db.ConnectDB()
	defer db.Close()

	grpcEnabled := os.Getenv("GRPC_ENABLED")

	cardService := factories.CardServiceFactory(db)
	paymentService := factories.PaymentServiceFactory(db)

	if grpcEnabled == "true" {
		runGRPC(db, cardService, paymentService)
	} else {
		runHTTP(db, cardService, paymentService)
	}
}

func runGRPC(db *pgxpool.Pool, cardService *card.CardService, paymentService *payment.PaymentService) {
	log.Println("[gRPC] Starting gRPC server...")

	grpcPort := os.Getenv("PORT")
	grpcServer := grpc.NewServer()
	pb.RegisterCardsServer(grpcServer, grpccard.NewCardServiceServer(db, cardService, paymentService))
	reflection.Register(grpcServer)
	lis, err := net.Listen("tcp", ":"+grpcPort)

	if err != nil {
		log.Fatal("[gRPC] Server closed unexpected", err.Error())
	}

	log.Println("[gRPC] Server started on port: " + grpcPort)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatal("[gRPC] Server closed unexpected")
	}
}

func runHTTP(db *pgxpool.Pool, cardService *card.CardService, paymentService *payment.PaymentService) {
	log.Println("[HTTP] Starting HTTP server...")
	appPort := os.Getenv("PORT")
	mux := api.NewApi(cardService, paymentService)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", appPort),
		Handler: mux,
	}

	log.Println("[HTTP] Server started on port " + appPort)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("[HTTP] Server closed unexpected")
	}
}
