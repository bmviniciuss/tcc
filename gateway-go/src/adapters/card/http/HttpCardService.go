package httpcardservice

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/bmviniciuss/gateway/src/core/card"
	"github.com/go-resty/resty/v2"
)

type HttpCardService struct {
	Client *resty.Client
}

func NewHttpCardService() *HttpCardService {
	return &HttpCardService{
		Client: resty.New(),
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

type CardRequestBody struct {
	CardholderName string `json:"cardholder_name"`
	IsCredit       bool   `json:"is_credit"`
	IsDebit        bool   `json:"is_debit"`
}

func (s *HttpCardService) CreateCard(input *card.CreateCardRequest) (*card.PresentationCard, error) {

	body := &CardRequestBody{
		CardholderName: input.CardholderName,
		IsCredit:       *input.IsCredit,
		IsDebit:        *input.IsDebit,
	}

	// postBody, _ := json.Marshal(body)
	// requestBody := bytes.NewBuffer(postBody)
	// resp, err := http.Post("http://0.0.0.0:5001/api/cards", "application/json", requestBody)

	// if err != nil {
	// 	log.Println("[HTTPCardService] Error while making request to crete a new card")
	// 	log.Println("[HTTPCardService] error: ", err)
	// 	return nil, errors.New("Unable to create card")
	// }

	// defer resp.Body.Close()

	// result := &CardResponse{}

	// respBody, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// err = json.Unmarshal(respBody, result)

	// if err != nil {
	// 	log.Println("[HTTPCardService] Error while parsing request body")
	// 	log.Println("[HTTPCardService] error: ", err)
	// 	return nil, errors.New("Unable to create card")
	// }
	host := os.Getenv("CARD_HOST")
	result := &CardResponse{}
	response, err := s.Client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Connection", "close").
		SetBody(body).
		// SetResult(result).
		Post(fmt.Sprintf("http://%s/api/cards", host))

	if err != nil {
		log.Println("[HTTPCardService] Error while making request to crete a new card")
		log.Println("[HTTPCardService] error: ", err)
		return nil, errors.New("Unable to create card")
	}

	err = json.Unmarshal(response.Body(), &result)

	if err != nil {
		log.Println("[HTTPCardService] Error while parsing request body")
		log.Println("[HTTPCardService] error: ", err)
		return nil, errors.New("Unable to create card")
	}

	presentationCard := &card.PresentationCard{
		Id:              result.Id,
		CardholderName:  result.CardholderName,
		Token:           result.Token,
		MaskedNumber:    result.MaskedNumber,
		ExpirationYear:  result.ExpirationYear,
		ExpirationMonth: result.ExpirationMonth,
		Active:          result.Active,
		IsCredit:        result.IsCredit,
		IsDebit:         result.IsDebit,
	}

	return presentationCard, nil

}
