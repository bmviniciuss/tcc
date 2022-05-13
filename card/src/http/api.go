package api

import (
	"time"

	postgrescardrepository "github.com/bmviniciuss/tcc/card/src/adapter/card"
	carddetailsgenerator "github.com/bmviniciuss/tcc/card/src/adapter/carddetails"
	"github.com/bmviniciuss/tcc/card/src/core/card"
	"github.com/bmviniciuss/tcc/card/src/core/encrypter"
	"github.com/bmviniciuss/tcc/card/src/http/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
)

func newCardService(db *sqlx.DB) *card.CardService {
	cardRepository := postgrescardrepository.NewPostgresCardRepository(db)
	encrypter := encrypter.NewEncrypter([]byte("gFvJR96@UXYrq_2m"))
	cardDetailsGenerator := carddetailsgenerator.NewCardDetailsGenerator()
	cardService := card.NewCardService(cardDetailsGenerator, encrypter, cardRepository)
	return cardService
}

func NewApi(db *sqlx.DB) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api", func(r chi.Router) {
		r.Route("/cards", func(r chi.Router) {
			handlers.NewCardsController(newCardService(db)).Route(r)
		})
	})

	return r
}
