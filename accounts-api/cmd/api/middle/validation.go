package middle

import (
	"accounts-api/internal/account"
	"accounts-api/internal/clients"
	"accounts-api/pkg/web"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	ErrMismatchedID  = errors.New("user id in token is different to user id in account")
	ErrInvalidUrlID  = errors.New("invalid id in url")
	ErrTokenRequired = errors.New("access token required")
	ErrInvalidToken  = errors.New("invalid token")
)

type Validation struct {
	accountService account.Service
}

func New(accountService account.Service) *Validation {
	return &Validation{
		accountService: accountService,
	}
}

// AccountUserIDMatchesAuthHeader valida que la domain.Account del accountID de la URL exista y tenga un campo UserID igual al UserID del token. AccountUserIDMatchesAuthHeader solo debería usarse si previamente se usa un AuthRequired.
// El UserID del Token fue guardado en la Key 'UserID' del gin.Context por el middle.AuthRequired.
func (m *Validation) AccountUserIDMatchesAuthHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accountID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Abort(ctx, 400, ErrInvalidUrlID)
			return
		}
		acc, err := m.accountService.GetByID(accountID)
		if err != nil {
			web.Abort(ctx, 404, account.ErrNotFound)
			return
		}
		userID := ctx.GetString("UserID")
		if userID != acc.UserID {
			web.Abort(ctx, 403, ErrMismatchedID)
			return
		}
		ctx.Next()
	}
}

// AuthRequired valida que haya un header 'Authorization'. Si lo hay, lo valida contra users-api (Keycloak).
// Setea en la ctx.Key 'UserID' el id del usuario.
func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("Authorization")
		if token == "" {
			web.Abort(ctx, 401, ErrTokenRequired)
			return
		}
		claims, err := clients.ValidateUserToken(token)
		if err != nil {
			web.Abort(ctx, 401, ErrInvalidToken)
			return
		}
		// Obtiene el campo 'sub' de las claims del JWT, que es el UserID, y lo setea en la key 'UserID'.
		ctx.Set("UserID", claims["sub"])

		ctx.Next()
	}
}
