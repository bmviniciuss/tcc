package api

import (
	"time"

	"github.com/bmviniciuss/tcc/card-payment/src/factories"
	paymenthandler "github.com/bmviniciuss/tcc/card-payment/src/http/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
)

func NewApi(db *sqlx.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api", func(r chi.Router) {
		r.Route("/payment", func(r chi.Router) {
			paymenthandler.NewPaymentController(factories.NewPaymentService(db)).Route(r)
		})
	})

	return r
}
