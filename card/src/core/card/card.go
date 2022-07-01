package card

import (
	"log"

	carddetails "github.com/bmviniciuss/tcc/card/src/core/cardDetails"
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

func (s *CardService) Generate(generateCardServiceInput *GenerateCardServiceInput) (*Card, error) {
	log.Println("[CardService] Generating card")
	log.Println("[CardService] Generating card details")
	cardDetails, err := s.cardDetailsGenerator.Generate()

	if err != nil {
		log.Println("[CardService] Error generating card details")
		return nil, err
	}

	log.Println("[CardService] Generating card expiration")
	year, month := getCardExpiration()
	log.Println("[CardService] Masking card PAN")
	MaskedNumber := MaskPANNumber(cardDetails.Number)

	generateCardInput := &GenerateCardRepoInput{
		Number:          cardDetails.Number,
		Cvv:             cardDetails.Cvv,
		CardholderName:  generateCardServiceInput.CardholderName,
		MaskedNumber:    MaskedNumber,
		Active:          true,
		ExpirationYear:  year,
		ExpirationMonth: month,
		IsCredit:        generateCardServiceInput.IsCredit,
		IsDebit:         generateCardServiceInput.IsDebit,
		Token:           generateToken(),
	}

	log.Println("[CardService] Saving card")
	card, err := s.cardRepository.Generate(generateCardInput)

	if err != nil {
		log.Println("[CardService] Error saving card", err)
		return nil, err
	}

	return card, nil
}

func (s *CardService) GetByToken(token string) (*Card, error) {
	log.Println("[CardService] Getting card by token")
	card, err := s.cardRepository.GetByToken(token)

	if err != nil {
		log.Println("[CardService] Error getting card by Token", err)
		return nil, err
	}

	return card, nil
}
