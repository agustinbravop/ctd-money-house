package middle

import (
	"errors"
	"github.com/gin-gonic/gin"
	"users-api/internal/auth"
	"users-api/pkg/web"
)

type AuthRequired struct {
	authService auth.Service
}

func New(authService auth.Service) *AuthRequired {
	return &AuthRequired{authService: authService}
}

func (m *AuthRequired) AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		if token == "" {
			web.Abort(ctx, 401, errors.New("access token required"))
			return
		}
		_, claims, err := m.authService.DecodeToken(token)
		if err != nil {
			web.Abort(ctx, 401, errors.New("invalid token"))
			return
		}

		// Obtiene el campo 'sub' del JWT Payload, que es el UserID de Keycloak, y lo pasa a string.
		ctx.Set("UserID", (*claims)["sub"].(string))
		ctx.Next()
	}
}
