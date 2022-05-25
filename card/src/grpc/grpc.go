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
	log.Println("[gRPC] GenerateCard called")

	input := card.GenerateCardServiceInput{
		CardholderName: in.GetCardholderName(),
		IsCredit:       in.GetIsCredit(),
		IsDebit:        in.GetIsDebit(),
	}

	log.Println("[gRPC] Calling card service to generate card")
	card, err := s.CardService.Generate(&input)

	if err != nil {
		log.Println("[gRPC] Error generating card: ", err)
		return nil, err
	}

	log.Println("[gRPC] Card generated. Returning results")

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

func (s *CardServiceServer) GetCardByToken(ctx context.Context, in *pb.GetCardByTokenRequest) (*pb.Card, error) {
	log.Println("[gRPC] GetCardByToken called")
	token := in.GetToken()
	log.Println("[gRPC] Token: ", token)

	log.Println("[gRPC] Calling card service to get card by token")
	card, err := s.CardService.GetByToken(token)

	if err != nil {
		log.Println("[gRPC] Error getting card by token: ", err)
		return nil, err
	}

	log.Println("[gRPC] Card found. Returning results")

	return &pb.Card{
		Id:              card.Id,
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
