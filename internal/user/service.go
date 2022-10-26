package user

import (
	"ctd-money-house/internal/domain"
	"ctd-money-house/pkg/utils"
	"database/sql"
	"errors"
)

var (
	ErrInternal    = errors.New("internal server error")
	ErrGettingUser = errors.New("error getting created user")
	ErrNotFound    = errors.New("user not found")
)

type Service interface {
	GetByID(id int) (domain.User, error)
	GetAll() ([]domain.User, error)
	Update(id int, user domain.User) (domain.User, error)
	Create(domain.User) (domain.User, error)
	Delete(id int) error
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
			return domain.User{}, ErrInternal
		}
	}

	newUser := builNewUser(u, user)
	err = s.r.Update(newUser)
	if err != nil {
		return domain.User{}, ErrInternal
	}

	return newUser, nil
}

func (s *service) Create(user domain.User) (domain.User, error) {
	user.Cvu = s.generateCvu()
	user.Alias = s.generateAlias()

	userID, err := s.r.Create(user)
	if err != nil {
		return domain.User{}, ErrInternal
	}

	userCreated, err := s.r.GetByID(userID)
	if err != nil {
		return domain.User{}, ErrGettingUser
	}

	return userCreated, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, ErrNotFound):
			return ErrNotFound
		default:
			return ErrInternal
		}
	}
	return nil
}

func (s *service) generateCvu() string {
	var cvu string
	for {
		cvu = utils.GenerateCvu()
		var fieldMap = map[string]interface{}{
			"cvu": cvu,
		}
		if s.r.ValidateCvuOrAlias(fieldMap) {
			break
		}
	}
	return cvu
}

func (s *service) generateAlias() string {
	var alias string
	for {
		alias = utils.GenerateAlias()
		var fieldMap = map[string]interface{}{
			"alias": alias,
		}
		if s.r.ValidateCvuOrAlias(fieldMap) {
			break
		}
	}
	return alias
}

func builNewUser(u, user domain.User) domain.User {
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
	return user
}
