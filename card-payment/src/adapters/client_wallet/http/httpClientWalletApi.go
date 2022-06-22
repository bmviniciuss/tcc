package httpclientwallet

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bmviniciuss/tcc/card-payment/src/core/payment"
	"github.com/go-resty/resty/v2"
)

type CreateTransactionPayload struct {
	ClientId             string  `json:"client_id"`
	Amount               float64 `json:"amount"`
	Type                 string  `json:"type"`
	TransactionServiceId string  `json:"transaction_service_id"`
	Service              string  `json:"service"`
	TransactionDate      string  `json:"transaction_date"`
}

type CreateTransactionResult struct {
	Id        string `json:"id"`
	CreatedAt string `json:"created_at"`
}

type HTTPClientWalletTransactionAPI struct {
}

func NewHTTPClientWalletTransactionAPI() *HTTPClientWalletTransactionAPI {
	return &HTTPClientWalletTransactionAPI{}
}

type Transaction struct {
	Id        string `json:"id"`
	CreatedAt string `json:"created_at"`
}

func (c *HTTPClientWalletTransactionAPI) Create(input *payment.ClientWalletTransaction) error {
	log.Println("[HTTPClientWalletTransactionAPI] Creating Trasanction on client wallet")
	client := resty.New()

	url := fmt.Sprintf("http://%s/api/transactions", os.Getenv("CLIENT_WALLET_HOST"))
	log.Printf("Creating a new transaction on %s", url)

	createTransactionPayload := CreateTransactionPayload{
		ClientId:             input.ClientId,
		Amount:               input.Amount,
		Type:                 input.Type,
		TransactionServiceId: input.TransactionServiceId,
		Service:              input.Service,
		TransactionDate:      input.TransactionDate.Format(time.RFC3339),
	}

	fmt.Println(createTransactionPayload)

	result := &CreateTransactionResult{}
	_, err := client.R().EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetBody(createTransactionPayload).
		SetResult(result).
		Post(url)

	if err != nil {
		log.Println("[HTTPClientWalletTransactionAPI] Error while making request to client wallet to create transaction")
		log.Println("[HTTPClientWalletTransactionAPI] error: ", err)
		return errors.New("Error while creating transaction on client wallet")
	}

	createdAt, err := time.Parse(time.RFC3339, result.CreatedAt)

	if err != nil {
		log.Println("[HTTPClientWalletTransactionAPI] Error while parsing created date: ", err)
		return err
	}

	input.Id = result.Id
	input.CreatedAt = createdAt

	return nil
}
