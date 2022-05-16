package card

import (
	"strings"
	"time"

	carddetails "github.com/bmviniciuss/tcc/card/src/core/cardDetails"
	"github.com/bmviniciuss/tcc/card/src/core/encrypter"
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

type CardRepository interface {
	Generate(generateCardDTO *GenerateCardRepoInput) (*Card, error)
	GetByToken(token string) (*Card, error)
}

type CardService struct {
	cardDetailsGenerator carddetails.GeneratorService
	encrypter            encrypter.Encrypter
	cardRepository       CardRepository
}

func NewCardService(cardDetailsGenerator carddetails.GeneratorService, encrypter encrypter.Encrypter, repository CardRepository) *CardService {
	return &CardService{
		cardDetailsGenerator: cardDetailsGenerator,
		encrypter:            encrypter,
		cardRepository:       repository,
	}
}

func maskPANNumber(pan string) string {
	return pan[:4] + strings.Repeat("*", len(pan)-8) + pan[len(pan)-4:]
}

func getCardExpiration() (int, int) {
	EXPIRATION_STEP_YEARS := 5 // in years
	now := time.Now()
	year := now.Year() + EXPIRATION_STEP_YEARS
	month := int(now.Month())
	return year, month
}

func (s *CardService) Generate(generateCardServiceInput *GenerateCardServiceInput) (*Card, error) {
	cardDetails, err := s.cardDetailsGenerator.Generate()

	if err != nil {
		return nil, err
	}

	year, month := getCardExpiration()
	MaskedNumber := maskPANNumber(cardDetails.Number)

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
	}

	enctypedCardNumber, err := s.encrypter.Encrypt([]byte(cardDetails.Number))

	if err != nil {
		return nil, err
	}

	generateCardInput.Token = string(enctypedCardNumber)

	card, err := s.cardRepository.Generate(generateCardInput)

	if err != nil {
		return nil, err
	}

	return card, nil
}

func (s *CardService) GetByToken(token string) (*Card, error) {
	pan, err := s.encrypter.Decrypt([]byte(token))

	if err != nil {
		return nil, err
	}

	card, err := s.cardRepository.GetByToken(string(pan))
	if err != nil {
		return nil, err
	}

	return card, nil
}
