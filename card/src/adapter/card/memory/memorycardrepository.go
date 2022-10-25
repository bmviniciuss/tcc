package memorycardrepository

import (
	"github.com/bmviniciuss/tcc/card/src/core/card"
	"github.com/google/uuid"
)

type memoryCardRepository struct {
}

func NewPostgresCardRepository() *memoryCardRepository {
	return &memoryCardRepository{}
}

func (r *memoryCardRepository) Generate(card *card.Card) error {
	card.Id = uuid.NewString()

	return nil
}

func (r *memoryCardRepository) GetByToken(token string) (*card.Card, error) {
	c := &card.Card{
		Id:              uuid.NewString(),
		Number:          "1234123412341234",
		Cvv:             "123",
		CardholderName:  "Vinicius Barbosa",
		Token:           uuid.NewString() + uuid.NewString(),
		MaskedNumber:    "1234********1234",
		ExpirationYear:  2022,
		ExpirationMonth: 2,
		Active:          true,
		IsCredit:        true,
		IsDebit:         true,
	}

	return c, nil
}
