package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/bmviniciuss/tcc/card/src/adapter/db"
	grpccard "github.com/bmviniciuss/tcc/card/src/grpc"
	"github.com/bmviniciuss/tcc/card/src/grpc/pb"
	api "github.com/bmviniciuss/tcc/card/src/http"

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

	grpcEnabled := os.Getenv("GRPC_ENABLED")

	if grpcEnabled == "true" {
		grpcPort := os.Getenv("PORT")
		grpcServer := grpc.NewServer()
		pb.RegisterCardsServer(grpcServer, grpccard.NewCardServiceServer(db))
		reflection.Register(grpcServer)

		lis, err := net.Listen("tcp", ":"+grpcPort)

		if err != nil {
			log.Fatal("Server closed unexpected", err.Error())
		}

		log.Println("GRPC server started on port: " + grpcPort)
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatal("Server closed unexpected")
		}
	} else {
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
}
