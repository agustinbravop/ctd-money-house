package user

import (
	"ctd-money-house/internal/domain"
	"database/sql"
	"errors"
)

// Errors
var (
	ErrNotFound     = errors.New("user not found")
	ErrBD           = errors.New("error")
)

type Service interface {
	GetByID(id int) (domain.User, error)
	GetAll() ([]domain.User, error)
	// Create(user domain.User) (domain.User, error)
	// Delete(id int) error
	Update(id int, user domain.User) (domain.User, error)
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

func (s *service) Update(id int, u domain.User) (domain.User, error) {
	user, err := s.r.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, ErrNotFound
		} else {
			return domain.User{}, ErrBD
		}
	}

	newUser := builNewUser(u, user)
	err = s.r.Update(newUser)
	if err != nil {
		return domain.User{}, ErrBD
	}

	return newUser, nil
}

func builNewUser(u, user domain.User) domain.User{
	if u.Name != "" {
		user.Name = u.Name
	}
	if u.LastName != "" {
		user.LastName = u.LastName
	}
	if u.Dni != "" {
		user.Dni = u.Dni
	}
	if u.Email != "" {
		user.Email = u.Email
	}
	if u.Telephone != "" {
		user.Telephone = u.Telephone
	}
	if u.Cvu != 0 {
		user.Cvu = u.Cvu
	}
	if u.Alias != "" {
		user.Alias = u.Alias
	}

	return user
}