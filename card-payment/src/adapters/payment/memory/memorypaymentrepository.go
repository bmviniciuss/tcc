package memorypaymentrepository

import (
	"github.com/google/uuid"

	"github.com/bmviniciuss/tcc/card-payment/src/core/payment"
)

type memoryPaymentRepository struct {
}

func NewPaymentRepository() *memoryPaymentRepository {
	return &memoryPaymentRepository{}
}

func (r *memoryPaymentRepository) Create(payment *payment.Payment) error {
	payment.Id = uuid.NewString()
	payment.Payable.Id = uuid.NewString()

	return nil
}

func (r *memoryPaymentRepository) GetPaymentsByClientId(input *payment.GetPaymentsByClientIdInput) ([]payment.Payment, error) {
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
