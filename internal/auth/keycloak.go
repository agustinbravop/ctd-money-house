package auth

import (
	"context"
	"ctd-money-house/internal/domain"
	"github.com/Nerzal/gocloak/v12"
)

type KeycloakClient interface {
	GetUserByID(userID string) (domain.User, error)
	LoginUser(email, password string) (*JWT, error)
	LogoutUser(refreshToken string) error
	RefreshToken(refreshToken string) (*JWT, error)
	CreateUser(user domain.User) (string, error)
	SetPassword(userID, password string) error
}

type JWT = gocloak.JWT

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

// GetUserByID solicita los datos del usuario con el userID dado.
// El userID es el string GUID autogenerado por Keycloak. Es distinto al el domain.User.ID.
func (k *keycloakClient) GetUserByID(userID string) (domain.User, error) {
	kcUser, err := k.client.GetUserByID(k.ctx, k.jwt.AccessToken, k.realm, userID)
	if err != nil {
		return domain.User{}, err
	}
	return toDomainUser(kcUser), nil
}

// CreateUser registra un nuevo usuario en Keycloak y devuelve su userID.
func (k *keycloakClient) CreateUser(user domain.User) (string, error) {
	kcUser := fromDomainUser(user)
	return k.client.CreateUser(k.ctx, k.jwt.AccessToken, k.realm, kcUser)
}

// LoginUser realiza el inicio de sesión en Keycloak y retorna el JWT de la sesión.
func (k *keycloakClient) LoginUser(email, password string) (*JWT, error) {
	return k.client.Login(k.ctx, k.clientID, k.clientSecret, k.realm, email, password)
}

// LogoutUser invalida el Refresh Token del usuario. Keycloak es incapaz de invalidar Access Tokens.
func (k *keycloakClient) LogoutUser(refreshToken string) error {
	return k.client.Logout(k.ctx, k.clientID, k.clientSecret, k.realm, refreshToken)
}

// RefreshToken solicita un JWT nuevo para el usuario. Keycloak genera un nuevo Access Token.
// Por seguridad usa Refresh Token Rotation y genera un nuevo Refresh Token, invalidando el anterior.
func (k *keycloakClient) RefreshToken(refreshToken string) (*JWT, error) {
	return k.client.RefreshToken(k.ctx, refreshToken, k.clientID, k.clientSecret, k.realm)
}

// SetPassword establece una nueva contraseña para el usuario con el userID dado.
func (k *keycloakClient) SetPassword(userID, password string) error {
	return k.client.SetPassword(k.ctx, k.jwt.AccessToken, userID, k.realm, password, false)
}

func toDomainUser(kcUser *gocloak.User) domain.User {
	return domain.User{
		Name:     *kcUser.FirstName,
		LastName: *kcUser.LastName,
		Email:    *kcUser.Email,
	}
}

func fromDomainUser(user domain.User) gocloak.User {
	userEnabled := true
	emailVerified := true
	return gocloak.User{
		Username:      &user.Email,
		EmailVerified: &emailVerified,
		FirstName:     &user.Name,
		LastName:      &user.LastName,
		Email:         &user.Email,
		Enabled:       &userEnabled,
	}
}
