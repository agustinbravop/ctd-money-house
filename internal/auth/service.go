package auth

import (
	"ctd-money-house/internal/domain"
	"github.com/Nerzal/gocloak/v12"
)

type Service interface {
	LoginUser(email, password string) (*gocloak.JWT, error)
	GetUserByKeycloakID(userID string) (domain.User, error)
}

type service struct {
	kc KeycloakClient
}

func NewAuthService(kc KeycloakClient) Service {
	return service{kc}
}

func (s service) LoginUser(email, password string) (*gocloak.JWT, error) {
	return s.kc.LoginUser(email, password)
}

func (s service) GetUserByKeycloakID(userID string) (domain.User, error) {
	return s.kc.GetUserByID(userID)
}
