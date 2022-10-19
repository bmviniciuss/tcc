package main

import (
	"fmt"
	"log"
	"math"
	"net"
	"net/http"
	"os"

	"github.com/bmviniciuss/tcc/card-payment/src/adapters/db"
	"github.com/bmviniciuss/tcc/card-payment/src/factories"
	grpcpaymentserver "github.com/bmviniciuss/tcc/card-payment/src/grpc"
	"github.com/bmviniciuss/tcc/card-payment/src/grpc/pb"
	api "github.com/bmviniciuss/tcc/card-payment/src/http"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// err := godotenv.Load()

	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	dbConn := db.ConnectDB()
	defer dbConn.Close()

	appPort := os.Getenv("PORT")
	grpcEnabled := os.Getenv("GRPC_ENABLED")

	if grpcEnabled == "true" {
		runGrpc(dbConn, appPort)
	} else {
		runHttp(dbConn, appPort)
	}
}

func runGrpc(db *pgxpool.Pool, appPort string) {
	log.Println("MAX: ", math.MaxInt32)

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

func runHttp(db *pgxpool.Pool, appPort string) {
	paymentService := factories.NewPaymentService(db)
	mux := api.NewApi(paymentService)

	server := http.Server{
		Addr:    ":" + appPort,
		Handler: mux,
	}

	log.Println("[HTTP] Server started on port " + appPort)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("[HTTP] Server closed unexpected")
	}
}
