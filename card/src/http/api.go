package api

import (
	"time"

	"github.com/bmviniciuss/tcc/card/src/core/card"
	"github.com/bmviniciuss/tcc/card/src/http/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewApi(cardService *card.CardService) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api", func(r chi.Router) {
		r.Route("/cards", func(r chi.Router) {
			handlers.NewCardsController(cardService).Route(r)
		})
	})

	return r
}
