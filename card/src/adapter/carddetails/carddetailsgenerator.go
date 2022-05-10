package carddetailsgenerator

import (
	"crypto/rand"
	"math/big"

	carddetails "github.com/bmviniciuss/tcc/card/src/core/cardDetails"
	"github.com/jaswdr/faker"
)

type carddetailsgenerator struct {
}

func NewCardDetailsGenerator() *carddetailsgenerator {
	return &carddetailsgenerator{}
}

func (g *carddetailsgenerator) Generate() (*carddetails.CardDetails, error) {
	pan := generatePAN()
	cvv, err := generateCVV()

	if err != nil {
		return nil, err
	}

	return &carddetails.CardDetails{
		Number: pan,
		Cvv:    cvv,
	}, nil
}

func generatePAN() string {
	faker := faker.New()
	return faker.Payment().CreditCardNumber()
}

func generateCVV() (string, error) {
	var cvv string

	for i := 0; i < 3; i++ {
		randomDigit, err := rand.Int(rand.Reader, big.NewInt(10))

		if err != nil {
			return "", err
		}

		cvv += randomDigit.String()

	}
	return cvv, nil
}
