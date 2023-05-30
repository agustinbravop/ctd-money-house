package domain

import (
	"time"
)

type Card struct {
	ID           string  `json:"id,omitempty" bson:"_id"`
	CardNumber   string  `json:"cardNumber"`
	ExpiryDate   string  `json:"expirationDate"`
	Owner        string  `json:"owner"`
	SecurityCode string  `json:"securityCode"`
	Brand        string  `json:"brand"`
	Amount       float64 `json:"amount"`
	IsBlocked    bool    `json:"isBlocked"`
}

func (c *Card) IsExpired() bool {
	expiryTime, _ := time.Parse("01/06", c.ExpiryDate)
	return time.Now().After(expiryTime)
}
