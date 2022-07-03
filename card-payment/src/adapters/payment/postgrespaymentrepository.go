package postgrespaymentrepository

import (
	"log"
	"time"

	"github.com/bmviniciuss/tcc/card-payment/src/core/payment"
	"github.com/jmoiron/sqlx"
)

type Payment struct {
	Id             string    `db:"id"`
	ClientId       string    `db:"client_id"`
	PaymentType    string    `db:"payment_type"`
	Amount         float64   `db:"amount"`
	CardholderName string    `db:"cardholder_name"`
	CardToken      string    `db:"card_token"`
	MaskedNumber   string    `db:"masked_number"`
	PaymentDate    time.Time `db:"payment_date"`
	CreatedAt      time.Time `db:"created_at"`
}

type Payable struct {
	Id          string    `db:"id"`
	ClientId    string    `db:"client_id"`
	PaymentId   string    `db:"payment_id"`
	PaymentDate time.Time `db:"payment_date"`
	Amount      float64   `db:"amount"`
}

func (Payment) TableName() string {
	return "payments"
}

type postgresPaymentRepository struct {
	Db sqlx.DB
}

func NewPostgresPaymentRepository(db *sqlx.DB) *postgresPaymentRepository {
	return &postgresPaymentRepository{
		Db: *db,
	}
}

func (r *postgresPaymentRepository) Create(payment *payment.Payment) error {
	tx, err := r.Db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if err := createPayment(tx, payment); err != nil {
		return err
	}

	if err := createPayable(tx, payment); err != nil {
		return err
	}

	return tx.Commit()
}

func createPayment(tx *sqlx.Tx, payment *payment.Payment) error {
	sql := `
		INSERT INTO cardpaymentms.payments 
		(client_id, payment_type, amount, cardholder_name, card_token, masked_number, payment_date) 
		VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id
	`
	err := tx.QueryRow(
		sql,
		payment.ClientId,
		payment.PaymentType,
		payment.Amount,
		payment.PaymentInfo.CardholderName,
		payment.PaymentInfo.CardToken,
		payment.PaymentInfo.MaskedNumber,
		payment.PaymentDate).Scan(&payment.Id)

	if err != nil {
		return err
	}

	payment.Payable.PaymentId = payment.Id

	return nil
}

func createPayable(tx *sqlx.Tx, payment *payment.Payment) error {
	sql := `
		INSERT INTO cardpaymentms.payables 
		(client_id, payment_id, payment_date, amount) 
		VALUES($1, $2, $3, $4) RETURNING id
	`
	err := tx.QueryRow(
		sql,
		payment.Payable.ClientId,
		payment.Id,
		payment.Payable.PaymentDate,
		payment.Amount).Scan(&payment.Payable.Id)

	if err != nil {
		return err
	}

	return nil
}

func (r *postgresPaymentRepository) GetPaymentsByClientId(input *payment.GetPaymentsByClientIdInput) ([]payment.Payment, error) {
	log.Println("PostgresRepo.GetPaymentsByClientId: Process started: ", input.ClientId)
	pp := []Payment{}
	res := []payment.Payment{}
	err := r.Db.Select(&pp, "SELECT * FROM cardpaymentms.payments WHERE client_id=$1", input.ClientId)

	if err != nil {
		log.Println("PostgresRepo.GetPaymentsByClientId: Error in query", err)
		return res, err // TODO: use generic error in the future
	}

	for _, p := range pp {
		res = append(res, payment.Payment{
			Id:          p.Id,
			ClientId:    p.ClientId,
			Amount:      p.Amount,
			PaymentType: p.PaymentType,
			PaymentInfo: payment.PaymentInfo{
				MaskedNumber: p.MaskedNumber,
			},
			PaymentDate: p.PaymentDate,
		})
	}

	return res, nil
}
