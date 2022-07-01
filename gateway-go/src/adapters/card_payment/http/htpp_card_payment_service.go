package httpcardpayment

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bmviniciuss/gateway/src/core/card_payment"
	"github.com/go-resty/resty/v2"
)

type HttpCardPaymentService struct {
	Client *resty.Client
}

func NewHttpCardPaymentService() *HttpCardPaymentService {
	return &HttpCardPaymentService{
		Client: resty.New().SetTimeout(20 * time.Second),
	}
}

type CardPaymentRequest struct {
	ClientId    string                 `json:"client_id"`
	Amount      float64                `json:"amount"`
	PaymentType string                 `json:"payment_type"`
	PaymentDate string                 `json:"payment_date"`
	PaymentInfo CardPaymentRequestInfo `json:"payment_info"`
}

type CardPaymentRequestInfo struct {
	CardToken string `json:"card_token"`
}

type CardPaymentResult struct {
	Id          string          `json:"id"`
	ClientId    string          `json:"client_id"`
	Amount      float64         `json:"amount"`
	PaymentType string          `json:"payment_type"`
	PaymentDate string          `json:"payment_date"`
	PaymentInfo CardPaymentInfo `json:"payment_info"`
}

type CardPaymentInfo struct {
	MaskedNumber string `json:"masked_number"`
}

func (h *HttpCardPaymentService) CreatePayment(payment *card_payment.CardPayment) error {
	result := &CardPaymentResult{}

	b := &CardPaymentRequest{
		ClientId:    payment.ClientId,
		Amount:      payment.Amount,
		PaymentType: payment.PaymentType,
		PaymentDate: payment.PaymentDate,
		PaymentInfo: CardPaymentRequestInfo{
			CardToken: payment.PaymentInfo.CardToken,
		},
	}

	// log.Printf("Calling HTTP Card Payment Microsservice with %+v\n", b)

	res, err := h.Client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Connection", "close").
		SetBody(b).
		Post(fmt.Sprintf("http://%s/api/payment", os.Getenv("CARD_PAYMENT_HOST")))

	if err != nil {
		log.Println("Error in request to process new card payment", err)
		return errors.New("An error occur while creating the card payment")
	}

	err = json.Unmarshal(res.Body(), &result)

	if err != nil {
		log.Println("Error while parsing request body")
		return errors.New("An error occur while creating the card payment")
	}

	if res.StatusCode() != http.StatusCreated {
		log.Println("The response was not expected", res.StatusCode(), string(res.Body()))
		return errors.New("An error occur while creating the card payment")
	}

	// log.Println("Card payment created")
	payment.Id = result.Id
	payment.PaymentInfo.MaskedNumber = result.PaymentInfo.MaskedNumber

	return nil

}
