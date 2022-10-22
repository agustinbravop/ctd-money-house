package handlers

import (
	"ctd-money-house/internal/auth"
	"ctd-money-house/pkg/web"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authHandler struct {
	s auth.Service
}

func NewAuthHandler(s auth.Service) *authHandler {
	return &authHandler{
		s: s,
	}
}

func (h *authHandler) Login() gin.HandlerFunc {
	type RequestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(ctx *gin.Context) {
		var body RequestBody
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("failed to parse json body"))
		}

		jwt, err := h.s.LoginUser(body.Email, body.Password)
		if err != nil {
			// TODO: con (err != nil) no se puede diferenciar si el error fue del usuario o del servidor.
			// TODO: sería mejor si pudieramos dar un mensaje de error más personalizado.
			web.Failure(ctx, http.StatusBadRequest, errors.New("wrong password or email"))
		}
		web.Success(ctx, http.StatusOK, jwt)
	}
}

// Logout recibe el Refresh Token a invalidar del usuario en el header Authorization.
// Una vez invalidado el Refresh Token, el usuario no puede obtener nuevos Access Tokens.
func (h *authHandler) Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		refreshToken := ctx.GetHeader("Authorization")
		if refreshToken == "" {
			web.Failure(ctx, http.StatusBadRequest, errors.New("missing Authorization header"))
		}
		err := h.s.LogoutUser(refreshToken)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, errors.New("something went wrong"))
		}
		web.Success(ctx, http.StatusOK, nil)
	}
}

// RefreshToken recibe el Refresh Token a utilizar en el header Authorization, y devuelve un nuevo JWT.
func (h *authHandler) RefreshToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		refreshToken := ctx.GetHeader("Authorization")
		if refreshToken == "" {
			web.Failure(ctx, http.StatusBadRequest, errors.New("missing Authorization header"))
		}
		jwt, err := h.s.RefreshToken(refreshToken)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, errors.New("something went wrong"))
		}
		web.Success(ctx, http.StatusOK, jwt)
	}
}
