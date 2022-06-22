package grpcclientwallet

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/bmviniciuss/tcc/card-payment/src/core/payment"
	"github.com/bmviniciuss/tcc/card-payment/src/grpc/pb"
	"google.golang.org/grpc"
)

type GRPCClientWalletAPI struct {
	Conn *grpc.ClientConn
}

func NewGRPCClientWalletAPI(Conn *grpc.ClientConn) *GRPCClientWalletAPI {
	return &GRPCClientWalletAPI{
		Conn: Conn,
	}
}

func (c *GRPCClientWalletAPI) Create(input *payment.ClientWalletTransaction) error {
	log.Println("[GRPCClientWalletAPI] Create")
	client := pb.NewClientWalletClient(c.Conn)

	transactionType, err := MapToTypeEnum(input.Type)

	if err != nil {
		log.Printf("Unsupported transaction type value = %s\n", input.Type)
		return errors.New("Unsupported transaction type provided")
	}

	serviceType, err := MapToServiceEnum(input.Service)
	if err != nil {
		log.Printf("Unsupported service type value = %s\n", input.Type)
		return errors.New("Unsupported service type provided")
	}

	serviceInput := &pb.CreateTransactionInput{
		ClientId:             input.ClientId,
		Amount:               input.Amount,
		Type:                 transactionType,
		TransactionServiceId: input.TransactionServiceId,
		TransactionDate:      input.TransactionDate.Format(time.RFC3339),
		Service:              serviceType,
	}

	resp, err := client.CreateTransaction(context.Background(), serviceInput)

	if err != nil {
		log.Println("[GRPCClientWalletAPI] Create error: ", err)
		return err
	}

	createdAt, err := time.Parse(time.RFC3339, resp.CreatedAt)

	if err != nil {
		log.Println("[GRPCClientWalletAPI] Error while parsing created date: ", err)
		return err
	}

	input.Id = resp.Id
	input.CreatedAt = createdAt
	return nil
}

func MapToTypeEnum(transactionType string) (pb.TransactionTypeEnum, error) {
	switch transactionType {
	case "CREDIT_CARD_PAYMENT":
		return pb.TransactionTypeEnum_CREDIT_CARD_PAYMENT, nil
	case "DEBIT_CARD_PAYMENT":
		return pb.TransactionTypeEnum_DEBIT_CARD_PAYMENT, nil
	case "WITHDRAWAL":
		return pb.TransactionTypeEnum_WITHDRAWAL, nil
	default:
		return -1, errors.New("Unsupported type for TransactionTypeEnum")
	}
}

func MapToServiceEnum(serviceType string) (pb.ServiceTypeEnum, error) {
	switch serviceType {
	case "CARD_PAYMENT":
		return pb.ServiceTypeEnum_CARD_PAYMENT, nil
	case "INTERNAL":
		return pb.ServiceTypeEnum_INTERNAL, nil
	default:
		return -1, errors.New("Unsupported type for ServiceTypeEnum")
	}
}
