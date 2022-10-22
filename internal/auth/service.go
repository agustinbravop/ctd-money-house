package auth

import (
	"ctd-money-house/internal/domain"
)

type Service interface {
	LoginUser(email, password string) (*JWT, error)
	GetUserByKeycloakID(userID string) (domain.User, error)
	LogoutUser(refreshToken string) error
	RefreshToken(refreshToken string) (*JWT, error)
}

type service struct {
	kc KeycloakClient
}

func NewAuthService(kc KeycloakClient) Service {
	return &service{kc}
}

func (s *service) LoginUser(email, password string) (*JWT, error) {
	return s.kc.LoginUser(email, password)
}

func (s *service) GetUserByKeycloakID(userID string) (domain.User, error) {
	return s.kc.GetUserByID(userID)
}

func (s *service) LogoutUser(refreshToken string) error {
	return s.kc.LogoutUser(refreshToken)
}

func (s *service) RefreshToken(refreshToken string) (*JWT, error) {
	return s.kc.RefreshToken(refreshToken)
}
