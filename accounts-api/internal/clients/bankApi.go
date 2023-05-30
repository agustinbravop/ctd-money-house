package clients

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

type ExtractionReq struct {
	CardNumber     string  `json:"card_number"`
	ExpiryDate     string  `json:"expiry_date"`
	Owner          string  `json:"owner"`
	SecurityCode   string  `json:"security_code"`
	Amount         float64 `json:"amount"`
	DestinationCvu string  `json:"destination_cvu"`
}

type ExtractionRes struct {
	Amount         float64 `json:"amount"`
	OriginCvu      string  `json:"origin_cvu"`
	DestinationCvu string  `json:"destination_cvu"`
}

var (
	ErrExtractionFromCard = errors.New("bank interface failed to extract from card")
)

func ExtractionFromCard(extractionReq ExtractionReq) (ExtractionRes, error) {
	host := os.Getenv("BANK_API_HOST")

	body, err := json.Marshal(extractionReq)
	if err != nil {
		return ExtractionRes{}, err
	}
	bodyReader := bytes.NewReader(body)
	req, err := http.NewRequest(http.MethodPost, host+"/api/v1/bank/extraction", bodyReader)
	if err != nil {
		return ExtractionRes{}, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil || res.StatusCode != 200 {
		return ExtractionRes{}, ErrExtractionFromCard
	}

	var extractionResponse ExtractionRes
	err = json.NewDecoder(res.Body).Decode(&extractionResponse)
	return extractionResponse, nil
}
