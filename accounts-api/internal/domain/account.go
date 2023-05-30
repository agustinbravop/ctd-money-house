package domain

type Account struct {
	ID     int     `json:"id"`
	Cvu    string  `json:"cvu"`
	Alias  string  `json:"alias"`
	Amount float64 `json:"amount"`
	UserID string  `json:"user_id"`
}
