package card

import (
	"bank-api/internal/domain"
)

var cards = []domain.Card{
	{
		CardNumber:   "4012000033330026",
		ExpiryDate:   "11/23",
		Owner:        "NORMAL OWNER",
		SecurityCode: "371",
		Brand:        "Visa",
		Amount:       54_350.57,
		IsBlocked:    false,
	},
	{
		CardNumber:   "374245455400126",
		ExpiryDate:   "05/23",
		Owner:        "UNLIMITED OWNER",
		SecurityCode: "537",
		Brand:        "Amex",
		Amount:       1_000_000_000.0,
		IsBlocked:    false,
	},
	{
		CardNumber:   "4205260000000005",
		ExpiryDate:   "04/23",
		Owner:        "EMPTY OWNER",
		SecurityCode: "837",
		Brand:        "Visa",
		Amount:       0.0,
		IsBlocked:    false,
	},
	{
		CardNumber:   "5425233430109903",
		ExpiryDate:   "12/04",
		Owner:        "EXPIRED OWNER",
		SecurityCode: "837",
		Brand:        "Mastercard",
		Amount:       987_654_321.0,
		IsBlocked:    false,
	},
	{
		CardNumber:   "5895626746595650",
		ExpiryDate:   "11/23",
		Owner:        "BLOCKED OWNER",
		SecurityCode: "837",
		Brand:        "Naranja",
		Amount:       1_200.0,
		IsBlocked:    true,
	},
}

type Repository interface {
	GetByCardNumber(cardNumber string) (domain.Card, error)
	GetByID(id string) (domain.Card, error)
	UpdateAmount(id string, newAmount float64) error
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetByID(id string) (domain.Card, error) {
	for _, card := range cards {
		if card.ID == id {
			return card, nil
		}
	}
	return domain.Card{}, ErrNotFound
}

func (r *repository) GetByCardNumber(cardNumber string) (domain.Card, error) {
	for _, card := range cards {
		if card.CardNumber == cardNumber {
			return card, nil
		}
	}
	return domain.Card{}, ErrNotFound
}

func (r *repository) UpdateAmount(id string, newAmount float64) error {
	for i := range cards {
		if cards[i].ID == id {
			cards[i].Amount = newAmount
			return nil
		}
	}
	return ErrNotFound
}
