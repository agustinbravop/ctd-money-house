package card

import (
	"accounts-api/internal/account"
	"accounts-api/internal/domain"
	"accounts-api/pkg/utils"
)

var (
	ErrInternal    = "internal server error"
	ErrGettingCard = "error getting just created card"
	ErrNotFound    = "card ID %d not found"
	ErrConflict    = "card is already associated to account: %d"
)

type Service interface {
	GetAllByAccountID(accountID int) ([]domain.Card, utils.ApiError)
	GetByID(id int) (domain.Card, utils.ApiError)
	Create(domain.Card) (domain.Card, utils.ApiError)
	Delete(idAccount, idCard int) utils.ApiError
}

type service struct {
	r  Repository
	as account.Service
}

func NewService(r Repository, as account.Service) Service {
	return &service{r, as}
}

func (s *service) GetAllByAccountID(accountID int) ([]domain.Card, utils.ApiError) {
	cards, err := s.r.GetAllByAccountID(accountID)
	if err != nil {
		return []domain.Card{}, err
	}
	return cards, nil
}

func (s *service) GetByID(id int) (domain.Card, utils.ApiError) {
	card, err := s.r.GetByID(id)
	if err != nil {
		return domain.Card{}, err
	}
	return card, nil
}

func (s *service) Create(card domain.Card) (domain.Card, utils.ApiError) {

	_, errGetByID := s.as.GetByID(card.AccountID)
	if errGetByID != nil {
		return domain.Card{}, utils.NewNotFoundError(errGetByID.Error(), errGetByID)
	}

	err := s.r.GetByCardNumber(card.CardNumber)
	if err != nil {
		return domain.Card{}, err
	}

	cardID, err := s.r.Create(card)
	if err != nil {
		return domain.Card{}, err
	}

	createdCard, err := s.r.GetByID(cardID)
	if err != nil {
		return domain.Card{}, err
	}

	return createdCard, nil
}

func (s *service) Delete(idAccount, idCard int) utils.ApiError {
	if err := s.r.Delete(idAccount, idCard); err != nil {
		return err
	}
	return nil
}
