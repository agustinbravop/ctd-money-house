package card

import (
	"bank-api/internal/domain"
	"errors"
)

var (
	ErrInsufficientFunds = errors.New("card funds are insufficient")
	ErrNotFound          = errors.New("card number not found")
	ErrDifferentFields   = errors.New("card data does not match")
	ErrExpiredCard       = errors.New("card has already expired")
	ErrBlockedCard       = errors.New("card is currently blocked")
	ErrInternal          = errors.New("internal server error")
)

type Extraction struct {
	Amount         float64 `json:"amount"`
	OriginCvu      string  `json:"origin_cvu"`
	DestinationCvu string  `json:"destination_cvu"`
}

type Service interface {
	GetByID(id string) (domain.Card, error)
	Extract(amount float64, destinationCvu string, card domain.Card) (Extraction, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id string) (domain.Card, error) {
	user, err := s.r.GetByID(id)
	if err != nil {
		return domain.Card{}, ErrNotFound
	}
	return user, nil
}

// Extract valida los datos de la domain.Card y decrementa extractionAmount a su amount si tiene suficiente saldo.
// Caso contrario, la operación no se realiza y retorna un error. Esto simula extraer dinero de la tarjeta.
func (s *service) Extract(extractionAmount float64, destinationCvu string, card domain.Card) (Extraction, error) {
	var extraction Extraction
	extraction.Amount = extractionAmount
	extraction.DestinationCvu = destinationCvu
	extraction.OriginCvu = "4036040911142433214202"

	if card.IsExpired() {
		return Extraction{}, ErrExpiredCard
	}
	// Casos hardcodeados para simular dinero ilimitado, insuficiente o bloqueada.
	switch {
	case card.Owner == "UNLIMITED OWNER":
		return extraction, nil
	case card.Owner == "EMPTY OWNER":
		return Extraction{}, ErrInsufficientFunds
	case card.Owner == "BLOCKED OWNER":
		return Extraction{}, ErrBlockedCard
	}

	return extraction, nil
}
