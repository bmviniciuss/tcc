package httpcardapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/bmviniciuss/tcc/card-payment/src/core/payment"
)

type HTTPCardAPI struct {
}

func NewHTTPCardAPI() *HTTPCardAPI {
	return &HTTPCardAPI{}
}

type CardResponse struct {
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

func (c *HTTPCardAPI) GetCardByToken(token string) (*payment.Card, error) {
	url := fmt.Sprintf("http://%s/api/cards?token=%s", os.Getenv("CARD_HOST"), token)
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var cardResponse CardResponse

	if err := json.NewDecoder(resp.Body).Decode(&cardResponse); err != nil {
		return nil, err
	}

	return &payment.Card{
		Id:              cardResponse.Id,
		CardholderName:  cardResponse.CardholderName,
		Token:           cardResponse.Token,
		MaskedNumber:    cardResponse.MaskedNumber,
		ExpirationYear:  cardResponse.ExpirationYear,
		ExpirationMonth: cardResponse.ExpirationMonth,
		Active:          cardResponse.Active,
		IsCredit:        cardResponse.IsCredit,
		IsDebit:         cardResponse.IsDebit,
	}, nil
}
