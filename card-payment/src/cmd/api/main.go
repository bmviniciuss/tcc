package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/bmviniciuss/tcc/card-payment/src/adapters/db"
	"github.com/bmviniciuss/tcc/card-payment/src/factories"
	grpcpaymentserver "github.com/bmviniciuss/tcc/card-payment/src/grpc"
	"github.com/bmviniciuss/tcc/card-payment/src/grpc/pb"
	api "github.com/bmviniciuss/tcc/card-payment/src/http"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := db.ConnectDB()
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	appPort := os.Getenv("PORT")
	grpcEnabled := os.Getenv("GRPC_ENABLED")

	if grpcEnabled == "true" {
		runGrpc(db, appPort)
	} else {
		runHttp(db, appPort)
	}
}

func runGrpc(db *sqlx.DB, appPort string) {
	gs := grpc.NewServer()
	paymentService := factories.NewPaymentService(db)
	pb.RegisterCardPaymentServer(gs, grpcpaymentserver.NewCardPaymentServer(paymentService))
	reflection.Register(gs)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", appPort))

	if err != nil {
		log.Fatal("[gRPC] Server closed unexpected", err.Error())
	}

	log.Printf("[gRPC] Server started on port: %s\n", appPort)

	if err = gs.Serve(lis); err != nil {
		log.Fatal("[gRPC] Server closed unexpected")
	}
}

func runHttp(db *sqlx.DB, appPort string) {

	mux := api.NewApi(db)
	server := http.Server{
		Addr:    ":" + appPort,
		Handler: mux,
	}

	log.Println("[HTTP] Server started on port " + appPort)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("[HTTP] Server closed unexpected")
	}
}
