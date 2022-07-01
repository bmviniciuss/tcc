package grpcpaymentserver

import (
	"context"
	"time"

	"github.com/bmviniciuss/tcc/card-payment/src/core/payment"
	"github.com/bmviniciuss/tcc/card-payment/src/grpc/pb"
)

type CardPaymentServer struct {
	pb.UnimplementedCardPaymentServer
	PaymentService payment.Service
}

func NewCardPaymentServer(service payment.Service) *CardPaymentServer {
	return &CardPaymentServer{
		PaymentService: service,
	}
}

func (s *CardPaymentServer) ProccessCardPayment(ctx context.Context, in *pb.ProcessCardPaymentInput) (*pb.Payment, error) {
	t, err := time.Parse(time.RFC3339, in.GetPaymentDate())
	if err != nil {
		return nil, err
	}
	p := int32(in.GetPaymentType())
	input := &payment.ProcessPaymentInput{
		ClientId:    in.GetClientId(),
		Amount:      float64(in.GetAmount()),
		PaymentType: pb.PaymentTypeEnum_name[p],
		PaymentDate: t,
		PaymentInfo: payment.PaymentInfoInput{
			CardToken: in.GetPaymentInfo().GetCardToken(),
		},
	}
	payment, err := s.PaymentService.Process(input)
	if err != nil {
		return nil, err
	}

	return &pb.Payment{
		Id:          payment.Id,
		ClientId:    payment.ClientId,
		Amount:      payment.Amount,
		PaymentType: payment.PaymentType,
		PaymentDate: payment.PaymentDate.Format(time.RFC3339),
		PaymentInfo: &pb.PaymentInfo{
			MaskedNumber: payment.PaymentInfo.MaskedNumber,
		},
	}, nil

}
