package postgrescardrepository

import (
	"github.com/bmviniciuss/tcc/card/src/core/card"
	"github.com/jmoiron/sqlx"
)

type PostgresCard struct {
	Id              string
	Number          string `db:"pan"`
	MaskedNumber    string `db:"masked_pan"`
	Cvv             string
	CardholderName  string `db:"cardholder_name"`
	Token           string
	ExpirationYear  int
	ExpirationMonth int
	Active          *bool
	IsCredit        *bool
	IsDebit         *bool
}

func (PostgresCard) TableName() string {
	return "cards"
}

type postgresCardRepository struct {
	Db sqlx.DB
}

func NewPostgresCardRepository(db *sqlx.DB) *postgresCardRepository {
	return &postgresCardRepository{
		Db: *db,
	}
}

func (r *postgresCardRepository) Generate(generateCardDTO *card.GenerateCardRepoInput) (*card.Card, error) {
	var id string

	pgCard := &PostgresCard{
		Number:          generateCardDTO.Number,
		Cvv:             generateCardDTO.Cvv,
		CardholderName:  generateCardDTO.CardholderName,
		Token:           generateCardDTO.Token,
		MaskedNumber:    generateCardDTO.MaskedNumber,
		ExpirationYear:  generateCardDTO.ExpirationYear,
		ExpirationMonth: generateCardDTO.ExpirationMonth,
		Active:          &generateCardDTO.Active,
		IsCredit:        &generateCardDTO.IsCredit,
		IsDebit:         &generateCardDTO.IsDebit,
	}

	insertSQL := "INSERT INTO public.cards (id, pan, masked_pan, cvv, cardholder_name, \"token\", expiration_year, expiration_month, active, is_debit, is_credit) VALUES(uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id"
	err := r.Db.QueryRow(insertSQL, pgCard.Number, pgCard.MaskedNumber, pgCard.Cvv, pgCard.CardholderName, pgCard.Token, pgCard.ExpirationYear, pgCard.ExpirationMonth, pgCard.Active, pgCard.IsDebit, pgCard.IsCredit).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &card.Card{
		Id:              id,
		Number:          pgCard.Number,
		Cvv:             pgCard.Cvv,
		CardholderName:  pgCard.CardholderName,
		Token:           pgCard.Token,
		MaskedNumber:    pgCard.MaskedNumber,
		ExpirationYear:  pgCard.ExpirationYear,
		ExpirationMonth: pgCard.ExpirationMonth,
		Active:          *pgCard.Active,
		IsCredit:        *pgCard.IsCredit,
		IsDebit:         *pgCard.IsDebit,
	}, nil
}

func (r *postgresCardRepository) GetByPan(pan string) (*card.Card, error) {
	var pgCard PostgresCard

	err := r.Db.Get(&pgCard, "SELECT * FROM public.cards c WHERE c.pan = $1", pan)

	if err != nil {
		return nil, err
	}

	return &card.Card{
		Id:              pgCard.Id,
		Number:          pgCard.Number,
		Cvv:             pgCard.Cvv,
		CardholderName:  pgCard.CardholderName,
		Token:           pgCard.Token,
		MaskedNumber:    pgCard.MaskedNumber,
		ExpirationYear:  pgCard.ExpirationYear,
		ExpirationMonth: pgCard.ExpirationMonth,
		Active:          *pgCard.Active,
		IsCredit:        *pgCard.IsCredit,
		IsDebit:         *pgCard.IsDebit,
	}, nil
}
