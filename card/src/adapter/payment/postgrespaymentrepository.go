package postgrespaymentrepository

import (
	"context"
	"time"

	"github.com/bmviniciuss/tcc/card/src/core/payment"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresPaymentAuthorization struct {
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

type postgresPaymentRepository struct {
	Db pgxpool.Pool
}

func NewPostgresPaymentRepository(db *pgxpool.Pool) *postgresPaymentRepository {
	return &postgresPaymentRepository{
		Db: *db,
	}
}

func (r *postgresPaymentRepository) Create(paymentAuthorization *payment.PaymentAuthorization) error {
	var id string
	var createAt time.Time

	insertSQL := "INSERT INTO cardms.payment_authorization" +
		"(id, amount, status, card_id, transaction_date)" +
		"VALUES(uuid_generate_v4(), $1, $2, $3, $4) returning id, created_at"

	err := r.Db.QueryRow(context.Background(), insertSQL, paymentAuthorization.Amount, paymentAuthorization.Status, paymentAuthorization.CardId, paymentAuthorization.TransactionDate).Scan(&id, &createAt)

	if err != nil {
		return err
	}

	paymentAuthorization.Id = id
	paymentAuthorization.CreatedAt = createAt

	return nil
}
