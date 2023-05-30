package auth

import (
	"context"
	"github.com/Nerzal/gocloak/v12"
	"github.com/golang-jwt/jwt/v4"
	"time"
	"users-api/internal/domain"
)

type KeycloakClient interface {
	GetUserByID(userID string) (domain.User, error)
	CreateUser(user domain.User) (string, error)
	LoginUser(email, password string) (*JWT, error)
	LogoutUser(refreshToken string) error
	UpdateUser(user domain.User) error
	DeleteUser(userID string) error
	SetPassword(userID, password string) error
	RefreshToken(refreshToken string) (*JWT, error)
	DecodeToken(accessToken string) (*jwt.Token, *jwt.MapClaims, error)
}

type JWT = gocloak.JWT

type keycloakClient struct {
	jwt            *gocloak.JWT
	client         *gocloak.GoCloak
	ctx            context.Context
	realm          string
	clientID       string
	clientSecret   string
	tokenExpiresAt int64
}

// NewKeycloakClient instancia un KeycloakClient. Puede fallar si la petición de LoginClient a Keycloak falla.
func NewKeycloakClient(url, clientID, clientSecret, realm string) (KeycloakClient, error) {
	k := &keycloakClient{
		client:       gocloak.NewClient(url),
		ctx:          context.Background(),
		realm:        realm,
		clientID:     clientID,
		clientSecret: clientSecret,
	}
	err := k.loginClient()
	if err != nil {
		return nil, err
	}
	return k, nil
}

// GetUserByID solicita los datos del usuario con el userID dado.
// El userID es el string GUID autogenerado por Keycloak. Es distinto al el domain.User.ID.
func (k *keycloakClient) GetUserByID(userID string) (domain.User, error) {
	k.getTokenIfExpired()
	kcUser, err := k.client.GetUserByID(k.ctx, k.jwt.AccessToken, k.realm, userID)
	if err != nil {
		return domain.User{}, err
	}
	return toDomainUser(kcUser), nil
}

// CreateUser registra un nuevo usuario en Keycloak y devuelve su userID.
func (k *keycloakClient) CreateUser(user domain.User) (string, error) {
	kcUser := fromDomainUser(user)
	k.getTokenIfExpired()
	return k.client.CreateUser(k.ctx, k.jwt.AccessToken, k.realm, kcUser)
}

// LoginUser realiza el inicio de sesión en Keycloak y retorna el JWT de la sesión.
func (k *keycloakClient) LoginUser(email, password string) (*JWT, error) {
	token, err := k.client.Login(k.ctx, k.clientID, k.clientSecret, k.realm, email, password)
	if err != nil {
		if gocloak.ParseAPIErrType(err) == gocloak.APIErrTypeInvalidGrant {
			return nil, errInvalidGrants
		} else {
			return nil, err
		}
	}
	return token, nil
}

// LogoutUser invalida el Refresh Token del usuario. Keycloak es incapaz de invalidar Access Tokens.
func (k *keycloakClient) LogoutUser(refreshToken string) error {
	return k.client.Logout(k.ctx, k.clientID, k.clientSecret, k.realm, refreshToken)
}

// RefreshToken solicita un JWT nuevo para el usuario. Keycloak genera un nuevo Access Token.
// Por seguridad se usa Refresh Token Rotation y genera un nuevo Refresh Token, invalidando el anterior.
func (k *keycloakClient) RefreshToken(refreshToken string) (*JWT, error) {
	return k.client.RefreshToken(k.ctx, refreshToken, k.clientID, k.clientSecret, k.realm)
}

// SetPassword establece una nueva contraseña para el usuario con el userID dado.
func (k *keycloakClient) SetPassword(userID, password string) error {
	k.getTokenIfExpired()
	return k.client.SetPassword(k.ctx, k.jwt.AccessToken, userID, k.realm, password, false)
}

// DeleteUser elimina un usuario de Keycloak según su userID.
func (k *keycloakClient) DeleteUser(userID string) error {
	k.getTokenIfExpired()
	return k.client.DeleteUser(k.ctx, k.jwt.AccessToken, k.realm, userID)
}

// UpdateUser actualiza los datos de un user dado.
func (k *keycloakClient) UpdateUser(user domain.User) error {
	k.getTokenIfExpired()
	return k.client.UpdateUser(k.ctx, k.jwt.AccessToken, k.realm, fromDomainUser(user))
}

// DecodeToken parsea y valida un accessToken emitido por Keycloak.
func (k *keycloakClient) DecodeToken(accessToken string) (*jwt.Token, *jwt.MapClaims, error) {
	return k.client.DecodeAccessToken(k.ctx, accessToken, k.realm)
}

// loginClient obtiene un nuevo JWT para el keycloakClient, y calcula el nuevo tokenExpiresAt.
func (k *keycloakClient) loginClient() error {
	token, err := k.client.LoginClient(k.ctx, k.clientID, k.clientSecret, k.realm)
	if err != nil {
		return err
	}
	k.jwt = token
	k.tokenExpiresAt = time.Now().Unix() + int64(token.ExpiresIn)
	return nil
}

// getTokenIfExpired valida (antes de cada petición) que el Access Token del KeycloakClient no esté vencido.
// Si el Access Token está vencido, realiza un loginClient del keycloakClient para obtener un token nuevo.
func (k *keycloakClient) getTokenIfExpired() {
	if k.tokenExpiresAt < time.Now().Unix() {
		err := k.loginClient()
		if err != nil {
			// Si el login del cliente falla, entonces la operación fracasará por 401 Unauthorized.
			return
		}
	}
}

func toDomainUser(kcUser *gocloak.User) domain.User {
	return domain.User{
		ID:       *kcUser.ID,
		Name:     *kcUser.FirstName,
		LastName: *kcUser.LastName,
		Email:    *kcUser.Email,
	}
}

func fromDomainUser(user domain.User) gocloak.User {
	isEnabled := true
	isEmailVerified := true
	return gocloak.User{
		ID:            &user.ID,
		Username:      &user.Email,
		EmailVerified: &isEmailVerified,
		FirstName:     &user.Name,
		LastName:      &user.LastName,
		Email:         &user.Email,
		Enabled:       &isEnabled,
	}
}
