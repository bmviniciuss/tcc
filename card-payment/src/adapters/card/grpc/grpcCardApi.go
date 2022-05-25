package grpccardapi

import (
	"context"
	"log"

	"github.com/bmviniciuss/tcc/card-payment/src/core/payment"
	"github.com/bmviniciuss/tcc/card-payment/src/grpc/pb"
	"google.golang.org/grpc"
)

type GRPCardAPI struct {
	Conn *grpc.ClientConn
}

func NewGRPCardAPI(Conn *grpc.ClientConn) *GRPCardAPI {

	return &GRPCardAPI{
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

func (c *GRPCardAPI) GetCardByToken(token string) (*payment.Card, error) {
	log.Println("[GRPCardAPI] GetCardByToken")

	client := pb.NewCardsClient(c.Conn)
	resp, err := client.GetCardByToken(context.Background(), &pb.GetCardByTokenRequest{Token: token})

	log.Printf("[GRPCardAPI] GetCardByToken resp: %v", resp)

	if err != nil {
		log.Println("[GRPCardAPI] GetCardByToken error: ", err)
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
