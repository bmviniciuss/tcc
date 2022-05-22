package grpccard

import (
	"context"
	"log"

	postgrescardrepository "github.com/bmviniciuss/tcc/card/src/adapter/card"
	carddetailsgenerator "github.com/bmviniciuss/tcc/card/src/adapter/carddetails"
	"github.com/bmviniciuss/tcc/card/src/core/card"
	"github.com/bmviniciuss/tcc/card/src/core/encrypter"
	"github.com/bmviniciuss/tcc/card/src/grpc/pb"
	"github.com/jmoiron/sqlx"
)

func newCardService(db *sqlx.DB) *card.CardService {
	cardRepository := postgrescardrepository.NewPostgresCardRepository(db)
	encrypter := encrypter.NewEncrypter([]byte("gFvJR96@UXYrq_2m"))
	cardDetailsGenerator := carddetailsgenerator.NewCardDetailsGenerator()
	cardService := card.NewCardService(cardDetailsGenerator, encrypter, cardRepository)
	return cardService
}

type CardServiceServer struct {
	pb.UnimplementedCardsServer
	CardService *card.CardService
}

func NewCardServiceServer(db *sqlx.DB) *CardServiceServer {
	cardService := newCardService(db)
	return &CardServiceServer{
		CardService: cardService,
	}
}

func (s *CardServiceServer) GenerateCard(ctx context.Context, in *pb.CreateCardRequest) (*pb.FullCard, error) {
	log.Printf("Received: %v %v %v\n", in.GetCardholderName(), in.GetIsCredit(), in.GetIsDebit())

	input := card.GenerateCardServiceInput{
		CardholderName: in.GetCardholderName(),
		IsCredit:       in.GetIsCredit(),
		IsDebit:        in.GetIsDebit(),
	}

	card, err := s.CardService.Generate(&input)

	if err != nil {
		return nil, err
	}

	return &pb.FullCard{
		Id:              card.Id,
		Number:          card.Number,
		Cvv:             card.Cvv,
		CardholderName:  card.CardholderName,
		Token:           card.Token,
		MaskedNumber:    card.MaskedNumber,
		ExpirationYear:  int64(card.ExpirationYear),
		ExpirationMonth: int64(card.ExpirationMonth),
		Active:          card.Active,
		IsCredit:        card.IsCredit,
		IsDebit:         card.IsDebit,
	}, nil
}
