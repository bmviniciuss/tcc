package api

import (
	"time"

	httpcardapi "github.com/bmviniciuss/tcc/card-payment/src/adapters/card"
	postgrespaymentrepository "github.com/bmviniciuss/tcc/card-payment/src/adapters/payment"
	"github.com/bmviniciuss/tcc/card-payment/src/core/payment"
	paymenthandler "github.com/bmviniciuss/tcc/card-payment/src/http/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
)

func NewPaymentService(db *sqlx.DB) payment.Service {
	repo := postgrespaymentrepository.NewPostgresPaymentRepository(db)
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
