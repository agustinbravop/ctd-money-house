package domain

type Card struct {
	ID             int    `json:"id,omitempty"`
	CardNumber     string `json:"card_number"`
	ExpirationDate string `json:"expiration_date"`
	Owner          string `json:"owner"`
	SecurityCode   string `json:"security_code"`
	Brand          string `json:"brand"`
	AccountID      int    `json:"account_id,omitempty"`
}
