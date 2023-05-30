package domain

type Transaction struct {
	ID              int     `json:"id"`
	Amount          float64 `json:"amount"`
	TransactionDate string  `json:"transaction_date"`
	Description     string  `json:"description"`
	OriginCvu       string  `json:"origin_cvu"`
	DestinationCvu  string  `json:"destination_cvu"`
	AccountID       int     `json:"account_id"`
	TransactionType string  `json:"transaction_type"`
}

type Filter struct {
	Type string `form:"type"`
	From string `form:"from,omitempty"`
	To   string `form:"to,omitempty"`
}
