package grpccardpayment

import (
	"context"
	"errors"
	"log"

	"github.com/bmviniciuss/gateway/src/core/card_payment"
	"github.com/bmviniciuss/gateway/src/grpc/pb"
	"google.golang.org/grpc"
)

type GRPCCardPaymentService struct {
	Client pb.CardPaymentClient
}

func NewGRPCCardPaymentService(conn *grpc.ClientConn) *GRPCCardPaymentService {
	return &GRPCCardPaymentService{
		Client: pb.NewCardPaymentClient(conn),
	}
}

func (h *GRPCCardPaymentService) CreatePayment(payment *card_payment.CardPayment) error {
	log.Println("[GRPCCardPaymentService] Process started")

	req := &pb.ProcessCardPaymentInput{
		ClientId:    payment.ClientId,
		PaymentType: pb.PaymentTypeEnum(pb.PaymentTypeEnum_value[payment.PaymentType]),
		PaymentDate: payment.PaymentDate,
		Amount:      payment.Amount,
		PaymentInfo: &pb.PaymentInfoInput{
			CardToken: payment.PaymentInfo.CardToken,
		},
	}

	resp, err := h.Client.ProccessCardPayment(context.Background(), req)

	if err != nil {
		log.Println("[gRPCCardPaymentService] Error:", err)
		return err
	}

	payment.Id = resp.GetId()
	payment.PaymentInfo.MaskedNumber = resp.GetPaymentInfo().GetMaskedNumber()
	return nil
}

func (h *GRPCCardPaymentService) GetPaymentsByClientId(clientId string) (*card_payment.CardPaymentsResponse, error) {
	return &card_payment.CardPaymentsResponse{}, errors.New("Not implemented")
}
