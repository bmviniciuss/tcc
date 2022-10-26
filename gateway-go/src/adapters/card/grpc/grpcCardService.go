package grpccardService

import (
	"context"
	"log"
	"os"

	"github.com/bmviniciuss/gateway/src/core/card"
	"github.com/bmviniciuss/gateway/src/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCardAPI struct {
	client pb.CardsClient
}

func NewGRPCardAPI() *GRPCardAPI {
	host := os.Getenv("CARD_HOST")
	grpcConn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal("Error connecting to card gRPC server")
	}

	return &GRPCardAPI{
		client: pb.NewCardsClient(grpcConn),
	}
}

func (c *GRPCardAPI) CreateCard(input *card.CreateCardRequest) (*card.PresentationCard, error) {
	req := &pb.CreateCardRequest{
		CardholderName: input.CardholderName,
		IsCredit:       *input.IsCredit,
		IsDebit:        *input.IsDebit,
	}

	resp, err := c.client.GenerateCard(context.Background(), req)

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
