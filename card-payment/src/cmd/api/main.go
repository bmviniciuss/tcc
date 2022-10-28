package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/bmviniciuss/tcc/card-payment/src/adapters/db"
	"github.com/bmviniciuss/tcc/card-payment/src/core/payment"
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

	paymentRepository := factories.PaymentRepositoryFactory(dbConn)
	cardApi := factories.NewCardApi()
	paymentService := factories.NewPaymentService(dbConn, cardApi, paymentRepository)

	if grpcEnabled == "true" {
		runGrpc(dbConn, appPort, paymentService)
	} else {
		runHttp(dbConn, appPort, paymentService)
	}
}

func runGrpc(db *pgxpool.Pool, appPort string, paymentService payment.Service) {
	gs := grpc.NewServer()

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

func runHttp(db *pgxpool.Pool, appPort string, paymentService payment.Service) {
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
