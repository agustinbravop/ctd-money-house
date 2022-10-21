package auth

import (
	"context"
	"ctd-money-house/internal/domain"
	"github.com/Nerzal/gocloak/v12"
)

type KeycloakClient interface {
	GetUserByID(userID string) (domain.User, error)
	LoginUser(email, password string) (*gocloak.JWT, error)
}

type keycloakClient struct {
	jwt          *gocloak.JWT
	client       *gocloak.GoCloak
	ctx          context.Context
	realm        string
	clientID     string
	clientSecret string
}

// NewKeycloakClient instancia un KeycloakClient. Puede fallar si la petición de LoginClient a Keycloak falla.
func NewKeycloakClient(url, clientID, clientSecret, realm string) (KeycloakClient, error) {
	client := gocloak.NewClient(url)
	ctx := context.Background()
	token, err := client.LoginClient(ctx, clientID, clientSecret, realm)
	if err != nil {
		return &keycloakClient{}, err
	}

	return &keycloakClient{
		jwt:          token,
		client:       client,
		clientSecret: clientSecret,
		clientID:     clientID,
		realm:        realm,
		ctx:          ctx,
	}, nil
}

// GetUserByID solicita a Keycloak los datos del usuario con el userID dado.
// El userID es el string GUID autogenerado por Keycloak. Es distinto al domain.User.ID y no están relacionados.
func (k *keycloakClient) GetUserByID(userID string) (domain.User, error) {
	kcUser, err := k.client.GetUserByID(k.ctx, k.jwt.AccessToken, k.realm, userID)
	if err != nil {
		return domain.User{}, err
	}
	return toDomainUser(kcUser), nil
}

// LoginUser realiza el inicio de sesión en Keycloak y retorna el JWT de la sesión.
func (k *keycloakClient) LoginUser(email, password string) (*gocloak.JWT, error) {
	return k.client.Login(k.ctx, k.clientID, k.clientSecret, k.realm, email, password)
}

func toDomainUser(kcUser *gocloak.User) domain.User {
	return domain.User{
		Name:     *kcUser.FirstName,
		LastName: *kcUser.LastName,
		Email:    *kcUser.Email,
	}
}
