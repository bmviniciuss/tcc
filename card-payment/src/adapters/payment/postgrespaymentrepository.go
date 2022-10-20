package postgrespaymentrepository

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/bmviniciuss/tcc/card-payment/src/core/payment"
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
	Db pgxpool.Pool
}

func NewPostgresPaymentRepository(db *pgxpool.Pool) *postgresPaymentRepository {
	return &postgresPaymentRepository{
		Db: *db,
	}
}

func (r *postgresPaymentRepository) Create(payment *payment.Payment) error {
	ctx := context.TODO()
	tx, err := r.Db.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return nil
	}
	defer tx.Rollback(context.TODO())

	err = createPayment(tx, payment)
	if err != nil {
		return err
	}

	err = createPayable(tx, payment)
	if err != nil {
		return err
	}

	err = tx.Commit(context.TODO())
	if err != nil {
		return err
	}

	return nil
}

func createPayment(tx pgx.Tx, payment *payment.Payment) error {
	sql := `
		INSERT INTO cardpaymentms.payments
		(id, client_id, payment_type, amount, cardholder_name, card_token, masked_number, payment_date, authorization_id)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id
	`
	err := tx.QueryRow(context.TODO(), sql,
		uuid.New().String(),
		payment.ClientId,
		payment.PaymentType,
		payment.Amount,
		payment.PaymentInfo.CardholderName,
		payment.PaymentInfo.CardToken,
		payment.PaymentInfo.MaskedNumber,
		payment.PaymentDate,
		payment.AuthorizationId).Scan(&payment.Id)

	if err != nil {
		log.Println("Error while creating payment = ", err.Error())
		return err
	}

	return nil
}

func createPayable(tx pgx.Tx, payment *payment.Payment) error {
	sql := `
		INSERT INTO cardpaymentms.payables
		(id, client_id, payment_id, payment_date, amount)
		VALUES($1, $2, $3, $4, $5) RETURNING id
	`
	err := tx.QueryRow(context.TODO(), sql,
		uuid.New().String(),
		payment.Payable.ClientId,
		payment.Id,
		payment.Payable.PaymentDate,
		payment.Amount,
	).Scan(&payment.Payable.Id)

	if err != nil {
		log.Println("Error while creating payable = ", err.Error())
		return err
	}

	return nil
}

func (r *postgresPaymentRepository) GetPaymentsByClientId(input *payment.GetPaymentsByClientIdInput) ([]payment.Payment, error) {
	//log.Println("PostgresRepo.GetPaymentsByClientId: Process started: ", input.ClientId)
	//pp := []Payment{}
	//res := []payment.Payment{}
	//err := r.Db.Select(&pp, "SELECT * FROM cardpaymentms.payments WHERE client_id=$1", input.ClientId)
	//
	//if err != nil {
	//	log.Println("PostgresRepo.GetPaymentsByClientId: Error in query", err)
	//	return res, err // TODO: use generic error in the future
	//}
	//
	//for _, p := range pp {
	//	res = append(res, payment.Payment{
	//		Id:          p.Id,
	//		ClientId:    p.ClientId,
	//		Amount:      p.Amount,
	//		PaymentType: p.PaymentType,
	//		PaymentInfo: payment.PaymentInfo{
	//			MaskedNumber: p.MaskedNumber,
	//		},
	//		PaymentDate: p.PaymentDate,
	//	})
	//}

	return []payment.Payment{}, nil
}
