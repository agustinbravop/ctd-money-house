package user

import (
	"errors"
	"strings"
	"users-api/internal/domain"
)

var (
	ErrInternal      = errors.New("internal server error")
	ErrNotFound      = errors.New("user not found")
	ErrMissingFields = errors.New("user is missing fields")
	ErrEmailInUse    = errors.New("email is already in use")
)

type Service interface {
	GetByID(id string) (domain.User, error)
	GetByEmail(email string) (domain.User, error)
	GetAll() ([]domain.User, error)
	Update(id string, user domain.User) (domain.User, error)
	Create(domain.User) (domain.User, error)
	Delete(id string) error
	ExistsByEmail(email string) bool
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id string) (domain.User, error) {
	user, err := s.r.GetByID(id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (s *service) GetByEmail(email string) (domain.User, error) {
	user, err := s.r.GetByEmail(email)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (s *service) GetAll() ([]domain.User, error) {
	users, err := s.r.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *service) Update(id string, user domain.User) (domain.User, error) {
	current, err := s.r.GetByID(id)
	if err != nil {
		return domain.User{}, err
	}
	newUser := buildNewUser(current, user)

	// Si el user tiene un email nuevo, se valida que no sea uno ya registrado.
	if current.Email != newUser.Email && s.r.ExistsByEmail(newUser.Email) {
		return domain.User{}, ErrEmailInUse
	}

	err = s.r.Update(newUser)
	if err != nil {
		return domain.User{}, ErrInternal
	}

	return newUser, nil
}

func (s *service) Create(user domain.User) (domain.User, error) {
	if isBlank(user.Name) || isBlank(user.LastName) || isBlank(user.Email) || isBlank(user.Dni) || isBlank(user.Telephone) {
		return domain.User{}, ErrMissingFields
	}

	userID, err := s.r.Create(user)
	if err != nil {
		return domain.User{}, ErrInternal
	}

	userCreated, err := s.r.GetByID(userID)
	if err != nil {
		return domain.User{}, ErrNotFound
	}

	return userCreated, nil
}

func (s *service) Delete(id string) error {
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

// ExistsByEmail retorna True si ya existe un domain.User con el email dado.
func (s *service) ExistsByEmail(email string) bool {
	return s.r.ExistsByEmail(email)
}

func isBlank(str string) bool {
	return strings.TrimSpace(str) == ""
}

func buildNewUser(current, new domain.User) domain.User {
	if new.Name != "" {
		current.Name = new.Name
	}
	if new.LastName != "" {
		current.LastName = new.LastName
	}
	if new.Dni != "" {
		current.Dni = new.Dni
	}
	if new.Email != "" {
		current.Email = new.Email
	}
	if new.Telephone != "" {
		current.Telephone = new.Telephone
	}
	return current
}
