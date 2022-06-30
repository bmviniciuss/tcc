package grpccardService

import (
	"context"
	"log"

	"github.com/bmviniciuss/gateway/src/core/card"
	"github.com/bmviniciuss/gateway/src/grpc/pb"
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

func (c *GRPCardAPI) CreateCard(input *card.CreateCardRequest) (*card.PresentationCard, error) {
	log.Println("[GRPCardService] CreateCard")
	client := pb.NewCardsClient(c.Conn)
	req := &pb.CreateCardRequest{
		CardholderName: input.CardholderName,
		IsCredit:       *input.IsCredit,
		IsDebit:        *input.IsDebit,
	}

	resp, err := client.GenerateCard(context.Background(), req)

	if err != nil {
		log.Println("[GRPCardService] GenerateCard error: ", err)
		return nil, err
	}

	return &card.PresentationCard{
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
