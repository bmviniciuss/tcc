package api

import (
	"log"
	"os"
	"time"

	httpcardapi "github.com/bmviniciuss/tcc/card-payment/src/adapters/card"
	grpccardapi "github.com/bmviniciuss/tcc/card-payment/src/adapters/card/grpc"
	postgrespaymentrepository "github.com/bmviniciuss/tcc/card-payment/src/adapters/payment"
	"github.com/bmviniciuss/tcc/card-payment/src/core/payment"
	paymenthandler "github.com/bmviniciuss/tcc/card-payment/src/http/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewPaymentService(db *sqlx.DB) payment.Service {
	repo := postgrespaymentrepository.NewPostgresPaymentRepository(db)

	if os.Getenv("GRPC_ENABLED") == "true" {
		host := os.Getenv("CARD_GRPC_HOST")
		grpcConn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))

		if err != nil {
			log.Fatal("Error connecting to grpc server")
		}

		grpcAPi := grpccardapi.NewGRPCCardClient(grpcConn)
		return payment.NewPaymentService(grpcAPi, repo)
	}

	cardAPI := httpcardapi.NewHTTPCardAPI()

	return payment.NewPaymentService(cardAPI, repo)
}

func NewApi(db *sqlx.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api", func(r chi.Router) {
		r.Route("/payment", func(r chi.Router) {
			paymenthandler.NewPaymentController(NewPaymentService(db)).Route(r)
		})
	})

	return r
}
