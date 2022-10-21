package grpccard

import (
	"context"
	"log"
	"time"

	"github.com/bmviniciuss/tcc/card/src/core/card"
	"github.com/bmviniciuss/tcc/card/src/core/payment"
	"github.com/bmviniciuss/tcc/card/src/factories"
	"github.com/bmviniciuss/tcc/card/src/grpc/pb"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CardServiceServer struct {
	pb.UnimplementedCardsServer
	CardService    *card.CardService
	paymentService *payment.PaymentService
}

func NewCardServiceServer(db *pgxpool.Pool) *CardServiceServer {
	cardService := factories.CardServiceFactory(db)
	paymentService := factories.PaymentServiceFactory(db)

	return &CardServiceServer{
		CardService:    cardService,
		paymentService: paymentService,
	}
}

func (s *CardServiceServer) GenerateCard(ctx context.Context, in *pb.CreateCardRequest) (*pb.FullCard, error) {

	input := card.GenerateCardServiceInput{
		CardholderName: in.GetCardholderName(),
		IsCredit:       in.GetIsCredit(),
		IsDebit:        in.GetIsDebit(),
	}

	card, err := s.CardService.Generate(&input)

	if err != nil {
		log.Println("[gRPC] Error generating card: ", err)
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

func (s *CardServiceServer) GetCardByToken(ctx context.Context, in *pb.GetCardByTokenRequest) (*pb.Card, error) {
	token := in.GetToken()

	card, err := s.CardService.GetByToken(token)

	if err != nil {
		log.Println("[gRPC] Error getting card by token: ", err)
		return nil, err
	}

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

func (s *CardServiceServer) AuthorizePayment(ctx context.Context, in *pb.AuhtorizePaymentRequest) (*pb.PaymentAuthorization, error) {
	f := "2006-01-02T15:04:05.000Z07:00"

	transctionDate, err := time.Parse(f, in.GetTrasanctionDate())

	if err != nil {
		return &pb.PaymentAuthorization{}, status.Error(codes.InvalidArgument, codes.InvalidArgument.String())
	}

	p := &payment.CreatePaymentAuthorization{
		Amount:          in.GetAmount(),
		CardToken:       in.GetCardToken(),
		PaymentType:     in.GetPaymentType(),
		TransactionDate: transctionDate,
	}

	authResult, err := s.paymentService.Authorize(p)

	if err != nil {
		return &pb.PaymentAuthorization{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.PaymentAuthorization{
		Id:              authResult.Id,
		Amount:          authResult.Amount,
		Status:          authResult.Status,
		TransactionDate: authResult.TransactionDate.UTC().Format(f),
		CreatedAt:       authResult.CreatedAt.UTC().Format(f),
	}, nil

}
