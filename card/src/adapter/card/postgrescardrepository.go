package postgrescardrepository

import (
	"context"

	"github.com/bmviniciuss/tcc/card/src/core/card"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresCard struct {
	Id              string `db:"id"`
	Number          string `db:"pan"`
	MaskedNumber    string `db:"masked_pan"`
	Cvv             string `db:"cvv"`
	CardholderName  string `db:"cardholder_name"`
	Token           string `db:"token"`
	ExpirationYear  int    `db:"expiration_year"`
	ExpirationMonth int    `db:"expiration_month"`
	Active          *bool  `db:"active"`
	IsCredit        *bool  `db:"is_credit"`
	IsDebit         *bool  `db:"is_debit"`
}

func (PostgresCard) TableName() string {
	return "cards"
}

type postgresCardRepository struct {
	Db pgxpool.Pool
}

func NewPostgresCardRepository(db *pgxpool.Pool) *postgresCardRepository {
	return &postgresCardRepository{
		Db: *db,
	}
}

func (r *postgresCardRepository) Generate(card *card.Card) error {
	var id string

	insertSQL := "INSERT INTO cardms.cards (id, pan, masked_pan, cvv, cardholder_name, \"token\", expiration_year, expiration_month, active, is_debit, is_credit) VALUES(uuid_generate_v4(), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id"
	err := r.Db.QueryRow(context.Background(), insertSQL, card.Number, card.MaskedNumber, card.Cvv, card.CardholderName, card.Token, card.ExpirationYear, card.ExpirationMonth, card.Active, card.IsDebit, card.IsCredit).Scan(&id)

	if err != nil {
		return err
	}

	card.Id = id

	return nil
}

func (r *postgresCardRepository) GetByToken(token string) (*card.Card, error) {
	// pgCard := PostgresCard{}
	return nil, nil
	// err := r.Db.Get(&pgCard, "select c.* from cardms.cards c where c.token=$1 LIMIT 1", token)

	// if err != nil {
	// 	return nil, err
	// }

	// return &card.Card{
	// 	Id:              pgCard.Id,
	// 	Number:          pgCard.Number,
	// 	Cvv:             pgCard.Cvv,
	// 	CardholderName:  pgCard.CardholderName,
	// 	Token:           pgCard.Token,
	// 	MaskedNumber:    pgCard.MaskedNumber,
	// 	ExpirationYear:  pgCard.ExpirationYear,
	// 	ExpirationMonth: pgCard.ExpirationMonth,
	// 	Active:          *pgCard.Active,
	// 	IsCredit:        *pgCard.IsCredit,
	// 	IsDebit:         *pgCard.IsDebit,
	// }, nil
}
