package auth

import (
	"ctd-money-house/internal/domain"
)

type Service interface {
	LoginUser(email, password string) (*JWT, error)
	LogoutUser(refreshToken string) error
	RefreshToken(refreshToken string) (*JWT, error)
	RegisterUser(user domain.User, password string) (domain.User, error)
}

type service struct {
	kc KeycloakClient
}

func NewAuthService(kc KeycloakClient) Service {
	return &service{kc}
}

func (s *service) RegisterUser(user domain.User, password string) (domain.User, error) {
	userID, err := s.kc.CreateUser(user)
	if err != nil {
		return domain.User{}, err
	}
	err = s.kc.SetPassword(userID, password)
	if err != nil {
		return domain.User{}, err
	}
	// TODO: Llamar al UserService para crear el usuario en el dominio (que debería cargar user.Alias y user.Cvu y devolverlo.
	// TODO: Si crear el usuario en el dominio falla, se puede usar el userID para eliminar al usuario de Keycloak (revertir la operación).
	return user, nil
}

func (s *service) LoginUser(email, password string) (*JWT, error) {
	return s.kc.LoginUser(email, password)
}

func (s *service) LogoutUser(refreshToken string) error {
	return s.kc.LogoutUser(refreshToken)
}

func (s *service) RefreshToken(refreshToken string) (*JWT, error) {
	return s.kc.RefreshToken(refreshToken)
}

func (s *service) getUserByKeycloakID(userID string) (domain.User, error) {
	return s.kc.GetUserByID(userID)
}
