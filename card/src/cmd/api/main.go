package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/bmviniciuss/tcc/card/src/adapter/db"
	"github.com/bmviniciuss/tcc/card/src/factories"
	grpccard "github.com/bmviniciuss/tcc/card/src/grpc"
	"github.com/bmviniciuss/tcc/card/src/grpc/pb"
	api "github.com/bmviniciuss/tcc/card/src/http"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("[Main] Error loading .env file")
	}

	db := db.ConnectDB()
	defer db.Close()

	grpcEnabled := os.Getenv("GRPC_ENABLED")

	if grpcEnabled == "true" {
		runGRPC(db)
	} else {
		runHTTP(db)
	}
}

func runGRPC(db *pgxpool.Pool) {
	log.Println("[gRPC] Starting gRPC server...")

	grpcPort := os.Getenv("PORT")
	var kaep = keepalive.EnforcementPolicy{
		MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
		PermitWithoutStream: true,            // Allow pings even when there are no active streams
	}

	var kasp = keepalive.ServerParameters{
		MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
		MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
		MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
		Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
		Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
	}
	grpcServer := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp))
	pb.RegisterCardsServer(grpcServer, grpccard.NewCardServiceServer(db))
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

func runHTTP(db *pgxpool.Pool) {
	log.Println("[HTTP] Starting HTTP server...")

	appPort := os.Getenv("PORT")
	cardService := factories.CardServiceFactory(db)
	paymentService := factories.PaymentServiceFactory(db)

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
