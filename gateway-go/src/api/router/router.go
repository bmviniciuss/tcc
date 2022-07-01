package router

import (
	"log"
	"os"

	grpccardservice "github.com/bmviniciuss/gateway/src/adapters/card/grpc"
	httpcardservice "github.com/bmviniciuss/gateway/src/adapters/card/http"
	httpcardpaymentservice "github.com/bmviniciuss/gateway/src/adapters/card_payment/http"
	"github.com/bmviniciuss/gateway/src/api/handler"
	m "github.com/bmviniciuss/gateway/src/api/middleware"
	"github.com/bmviniciuss/gateway/src/core/card"
	"github.com/bmviniciuss/gateway/src/core/card_payment"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetRouter() *chi.Mux {
	r := chi.NewRouter()

	// r.Use(cm.RequestID)
	// r.Use(cm.RealIP)
	// r.Use(cm.Logger)
	// r.Use(cm.Recoverer)
	// r.Use(cm.Timeout(5 * time.Second))
	r.Use(m.Json)

	buildTree(r)

	return r
}

func getCardService() card.Service {
	if os.Getenv("GRPC_ENABLED") == "true" {
		log.Println("Creating a gRPC card API")
		host := os.Getenv("CARD_HOST")
		grpcConn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err != nil {
			log.Fatal("Error connecting to card gRPC server")
		}
		return grpccardservice.NewGRPCardAPI(grpcConn)
	}

	log.Println("Creating a HTTP card API")
	return httpcardservice.NewHttpCardService()
}

func getCardPaymentService() card_payment.Service {
	if os.Getenv("GRPC_ENABLED") == "true" {
		log.Println("Creating GRPCCardPaymentService")
		log.Fatalln("Not implemented")
	}
	log.Println("Creating HttpCardPaymentService")
	return httpcardpaymentservice.NewHttpCardPaymentService()
}

func buildTree(r *chi.Mux) {
	s := getCardService()
	cps := getCardPaymentService()

	r.Route("/api", func(r chi.Router) {
		r.Route("/cards", func(r chi.Router) {
			handler.MakeCardHandlers(r, s)
		})
		r.Route("/payments", func(r chi.Router) {
			handler.MakePaymentHandlers(r, cps)
		})
	})
}
