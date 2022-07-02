package grpcpaymentserver

import (
	"context"
	"time"

	"github.com/bmviniciuss/tcc/card-payment/src/core/payment"
	"github.com/bmviniciuss/tcc/card-payment/src/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *CardPaymentServer) GetPaymentsByClientId(ctx context.Context, in *pb.GetPaymentsByClientIdRequest) (*pb.PaymentsResults, error) {
	id := in.GetClientId()
	if id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "The field 'client_id' must be provided.")
	}

	input := &payment.GetPaymentsByClientIdInput{
		ClientId: in.GetClientId(),
	}

	pp, err := s.PaymentService.GetPaymentsByClientId(input)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "An error occur while searching for the payments")
	}

	c := []*pb.Payment{}
	for _, payment := range pp {
		p := &pb.Payment{
			Id:          payment.Id,
			ClientId:    payment.ClientId,
			PaymentType: payment.PaymentType,
			PaymentDate: payment.PaymentDate.Format(time.RFC3339),
			Amount:      payment.Amount,
			PaymentInfo: &pb.PaymentInfo{
				MaskedNumber: payment.PaymentInfo.MaskedNumber,
			},
		}
		c = append(c, p)
	}

	res := pb.PaymentsResults{
		Content: c,
	}

	return &res, nil
}
