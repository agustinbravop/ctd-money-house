package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"users-api/internal/domain"
	"users-api/internal/user"
)

var (
	ErrInternal      = errors.New("internal server error")
	ErrMissingFields = errors.New("user is missing fields")
	ErrEmailInUse    = errors.New("email is already in use")
	ErrWrongPassword = errors.New("wrong password")
	ErrWrongEmail    = errors.New("email is not registered")
	errInvalidGrants = errors.New("invalid grants")
)

type Service interface {
	RegisterUser(user domain.User, password string) (domain.User, error)
	LoginUser(email, password string) (*JWT, error)
	LogoutUser(refreshToken string) error
	UpdateUser(user domain.User, password string) (domain.User, error)
	SetPassword(userID, newPassword string) error
	RefreshToken(refreshToken string) (*JWT, error)
	DecodeToken(accessToken string) (*jwt.Token, *jwt.MapClaims, error)
}

type service struct {
	kc          KeycloakClient
	userService user.Service
}

func NewAuthService(kc KeycloakClient, userService user.Service) Service {
	return &service{kc, userService}
}

func (s *service) RegisterUser(u domain.User, password string) (domain.User, error) {
	if isBlank(u.Name) || isBlank(u.LastName) || isBlank(u.Email) || isBlank(password) {
		return domain.User{}, ErrMissingFields
	}
	// Validar que el email no esté ocupado.
	if s.userService.ExistsByEmail(u.Email) {
		return domain.User{}, ErrEmailInUse
	}
	// Se registra el user en Keycloak.
	userID, err := s.kc.CreateUser(u)
	if err != nil {
		return domain.User{}, ErrInternal
	}

	err = s.kc.SetPassword(userID, password)
	if err != nil {
		return domain.User{}, ErrInternal
	}

	// Ahora se crea el user en nuestra base de datos.
	u.ID = userID
	createdUser, err := s.userService.Create(u)
	if err != nil {
		// Pseudo-rollback para eliminar el user creado en Keycloak.
		_ = s.kc.DeleteUser(userID)
		if errors.Is(err, user.ErrMissingFields) {
			return domain.User{}, ErrMissingFields
		} else {
			return domain.User{}, ErrInternal
		}
	}
	return createdUser, nil
}

func (s *service) LoginUser(email, password string) (*JWT, error) {
	if isBlank(email) || isBlank(password) {
		return nil, ErrMissingFields
	}

	token, err := s.kc.LoginUser(email, password)
	if err != nil {
		if errors.Is(err, errInvalidGrants) {
			if s.userService.ExistsByEmail(email) {
				return nil, ErrWrongPassword
			} else {
				return nil, ErrWrongEmail
			}
		} else {
			return nil, ErrInternal
		}
	}
	return token, nil
}

func (s *service) LogoutUser(refreshToken string) error {
	err := s.kc.LogoutUser(refreshToken)
	if err != nil {
		return ErrInternal
	}
	return nil
}

func (s *service) UpdateUser(user domain.User, password string) (domain.User, error) {
	// Primero se updatean los datos en la DB de users.Service.
	newUser, err := s.userService.Update(user.ID, user)
	if err != nil {
		return domain.User{}, err
	}
	// Luego se updatean en Keycloak.
	err = s.kc.UpdateUser(newUser)
	if err != nil {
		return domain.User{}, ErrInternal
	}
	// Si user tiene una password nueva, se la setea en Keycloak.
	if password != "" {
		if err = s.kc.SetPassword(newUser.ID, password); err != nil {
			return domain.User{}, ErrInternal
		}
	}
	return newUser, nil
}

func (s *service) RefreshToken(refreshToken string) (*JWT, error) {
	token, err := s.kc.RefreshToken(refreshToken)
	if err != nil {
		return nil, ErrInternal
	}
	return token, nil
}

func (s *service) SetPassword(userID, newPassword string) error {
	err := s.kc.SetPassword(userID, newPassword)
	if err != nil {
		return ErrInternal
	}
	return nil
}

// DecodeToken por dentro ya tiene en cuenta que el accessToken viene con el 'Bearer ' del HTTP Header 'Authentication'.
func (s *service) DecodeToken(accessToken string) (*jwt.Token, *jwt.MapClaims, error) {
	token, claims, err := s.kc.DecodeToken(accessToken)
	if err != nil {
		return nil, nil, ErrInternal
	}
	return token, claims, nil
}

func (s *service) getUserByID(userID string) (domain.User, error) {
	u, err := s.kc.GetUserByID(userID)
	if err != nil {
		return domain.User{}, ErrInternal
	}
	return u, nil
}

func isBlank(str string) bool {
	return strings.TrimSpace(str) == ""
}
