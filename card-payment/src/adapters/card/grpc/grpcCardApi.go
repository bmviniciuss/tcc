package grpccardapi

import (
	"context"
	"log"

	"github.com/bmviniciuss/tcc/card-payment/src/adapters/grpc/pb"
	"github.com/bmviniciuss/tcc/card-payment/src/core/payment"
	"google.golang.org/grpc"
)

type GRPCCardClient struct {
	Conn *grpc.ClientConn
}

func NewGRPCCardClient(Conn *grpc.ClientConn) *GRPCCardClient {

	return &GRPCCardClient{
		Conn: Conn,
	}
}

type CardResponse struct {
	Data Card `json:"data"`
}

type Card struct {
	Id              string `json:"id"`
	CardholderName  string `json:"cardholder_name"`
	Token           string `json:"token"`
	MaskedNumber    string `json:"masked_number"`
	ExpirationYear  int    `json:"expiration_year"`
	ExpirationMonth int    `json:"expiration_month"`
	Active          bool   `json:"active"`
	IsCredit        bool   `json:"is_credit"`
	IsDebit         bool   `json:"is_debit"`
}

func (c *GRPCCardClient) GetCardByToken(token string) (*payment.Card, error) {
	log.Println("[GRPCCardClient] GetCardByToken")

	client := pb.NewCardsClient(c.Conn)
	resp, err := client.GetCardByToken(context.Background(), &pb.GetCardByTokenRequest{Token: token})

	log.Printf("[GRPCCardClient] GetCardByToken resp: %v", resp)

	if err != nil {
		log.Println("[GRPCCardClient] GetCardByToken error: ", err)
		return nil, err
	}

	return &payment.Card{
		Id:              resp.GetId(),
		CardholderName:  resp.GetCardholderName(),
		Token:           resp.GetToken(),
		MaskedNumber:    resp.GetMaskedNumber(),
		ExpirationYear:  int(resp.GetExpirationYear()),
		ExpirationMonth: int(resp.GetExpirationMonth()),
		Active:          resp.GetActive(),
		IsCredit:        resp.GetIsCredit(),
		IsDebit:         resp.GetIsDebit(),
	}, nil
}
