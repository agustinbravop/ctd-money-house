package user

import "github.com/agustinbravop/ctd-money-house/internal/domain"

type Service interface {
	GetByID(id int) (domain.User, error)
	GetAll() ([]domain.User, error)
	// Create(p domain.User) (domain.User, error)
	// Delete(id int) error
	// Update(id int, p domain.User) (domain.User, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domain.User, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return domain.User{}, err
	}
	return d, nil
}

func (s *service) GetAll() ([]domain.User, error) {
	users, err := s.r.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
