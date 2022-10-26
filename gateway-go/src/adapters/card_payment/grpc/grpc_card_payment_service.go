package grpccardpayment

import (
	"context"
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
	result := &card_payment.CardPaymentsResponse{}
	input := &pb.GetPaymentsByClientIdRequest{
		ClientId: clientId,
	}

	res, err := h.Client.GetPaymentsByClientId(context.Background(), input)

	if err != nil {
		log.Println("[gRPCCardPaymentService.GetPaymentsByClientId] Error:", err)
		return result, err
	}

	for _, p := range res.Content {
		result.Content = append(result.Content, card_payment.CardPayment{
			Id:          p.GetId(),
			ClientId:    p.GetClientId(),
			Amount:      p.GetAmount(),
			PaymentType: p.GetPaymentType(),
			PaymentDate: p.GetPaymentDate(),
			PaymentInfo: card_payment.CardPaymentInfo{
				MaskedNumber: p.GetPaymentInfo().GetMaskedNumber(),
			},
		})
	}

	return result, nil
}
