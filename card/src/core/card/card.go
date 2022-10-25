package card

import (
	"log"

	carddetails "github.com/bmviniciuss/tcc/card/src/core/cardDetails"
	"github.com/google/uuid"
)

type Card struct {
	Id              string
	Number          string `json:"-"`
	Cvv             string `json:"-"`
	CardholderName  string
	Token           string
	MaskedNumber    string
	ExpirationYear  int
	ExpirationMonth int
	Active          bool
	IsCredit        bool
	IsDebit         bool
}

type GenerateCardServiceInput struct {
	CardholderName string
	IsCredit       bool
	IsDebit        bool
}

type Service interface {
	Generate(generateCardDTO *GenerateCardServiceInput) (*Card, error)
	GetByToken(token string) (*Card, error)
}

type GenerateCardRepoInput struct {
	Number          string
	Cvv             string
	CardholderName  string
	Token           string
	MaskedNumber    string
	ExpirationYear  int
	ExpirationMonth int
	Active          bool
	IsCredit        bool
	IsDebit         bool
}

type CardService struct {
	cardDetailsGenerator carddetails.GeneratorService
	cardRepository       CardRepository
}

func NewCardService(cardDetailsGenerator carddetails.GeneratorService, repository CardRepository) *CardService {
	return &CardService{
		cardDetailsGenerator: cardDetailsGenerator,
		cardRepository:       repository,
	}
}

func (s *CardService) Generate(generateCardDTO *GenerateCardServiceInput) (*Card, error) {
	cardDetails, err := s.cardDetailsGenerator.Generate()

	if err != nil {
		log.Println("[CardService] Error generating card details")
		return nil, err
	}

	year, month := getCardExpiration()
	MaskedNumber := MaskPANNumber(cardDetails.Number)

	card := &Card{
		Id:              uuid.New().String(),
		Number:          cardDetails.Number,
		Cvv:             cardDetails.Cvv,
		CardholderName:  generateCardDTO.CardholderName,
		Token:           generateToken(),
		MaskedNumber:    MaskedNumber,
		ExpirationYear:  year,
		ExpirationMonth: month,
		Active:          true,
		IsCredit:        generateCardDTO.IsCredit,
		IsDebit:         generateCardDTO.IsDebit,
	}

	err = s.cardRepository.Generate(card)

	if err != nil {
		log.Println("[CardService] Error saving card", err)
		return nil, err
	}

	return card, nil
}

func (s *CardService) GetByToken(token string) (*Card, error) {
	card, err := s.cardRepository.GetByToken(token)

	if err != nil {
		log.Println("[CardService] Error getting card by Token", err)
		return nil, err
	}

	return card, nil
}
