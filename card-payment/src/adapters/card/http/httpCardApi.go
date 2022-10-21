package httpcardapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bmviniciuss/tcc/card-payment/src/core/payment"
	"github.com/go-resty/resty/v2"
)

type HTTPCardAPI struct {
	client *resty.Client
}

func NewHTTPCardAPI() *HTTPCardAPI {
	r := resty.New()
	// t := http.DefaultTransport.(*http.Transport).Clone()

	// t.MaxIdleConnsPerHost = 100

	// r.SetTransport(t)
	return &HTTPCardAPI{
		client: r,
	}
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

type PaymentAuthorizationRequest struct {
	Amount          float64 `json:"amount"`
	CardToken       string  `json:"card_token"`
	PaymentType     string  `json:"payment_type"`
	TransactionDate string  `json:"transaction_date"`
}

type PaymentAuthorizationResponse struct {
	Id              string  `json:"id"`
	Amount          float64 `json:"amount"`
	Status          string  `json:"status"`
	TransactionDate string  `json:"transaction_date"`
	CreateAt        string  `json:"create_at"`
}

func (c *HTTPCardAPI) GetCardByToken(token string) (*payment.Card, error) {
	url := fmt.Sprintf("http://%s/api/cards?token=%s", os.Getenv("CARD_HOST"), token)
	resp, err := http.Get(url)

	if err != nil {
		log.Println("Error = ", err.Error())
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

func (c *HTTPCardAPI) AuthorizePayment(input *payment.PaymentAuthorizationInput) (*payment.PaymentAuthorization, error) {
	fmt.Println("[HTTPCardAPI] AuthorizePayment")
	url := fmt.Sprintf("http://%s/api/payment/authorize", os.Getenv("CARD_HOST"))

	requestBody := &PaymentAuthorizationRequest{
		Amount:          input.Amount,
		CardToken:       input.CardToken,
		PaymentType:     input.PaymentType,
		TransactionDate: input.TransactionDate,
	}

	response := &PaymentAuthorizationResponse{}

	_, err := c.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(requestBody).
		SetResult(response).
		Post(url)

	fmt.Printf("%v\n", response)

	if err != nil {
		fmt.Println("Error = ", err.Error())

		return nil, err
	}

	return &payment.PaymentAuthorization{
		Id:              response.Id,
		Amount:          response.Amount,
		Status:          response.Status,
		TransactionDate: response.TransactionDate,
		CreateAt:        response.CreateAt,
	}, nil
}
