package handlers

import (
	"ctd-money-house/internal/auth"
	"ctd-money-house/internal/domain"
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

// Register crea un nuevo usuario, y retorna el usuario creado.
func (h *authHandler) Register() gin.HandlerFunc {
	type RequestBody struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		Dni       string `json:"dni"`
		Telephone string `json:"telephone"`
	}

	return func(ctx *gin.Context) {
		var body RequestBody
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("failed to parse json body"))
		}
		user := domain.User{
			Name:      body.FirstName,
			LastName:  body.LastName,
			Dni:       body.Dni,
			Email:     body.Email,
			Telephone: body.Telephone,
		}
		user, err = h.s.RegisterUser(user, body.Password)
		if err != nil {
			// TODO: personalizar mensaje de error si es que el email dado ya existe (debe ser único).
			web.Failure(ctx, http.StatusInternalServerError, errors.New("something went wrong"))
		}
		web.Success(ctx, http.StatusCreated, user)
	}
}

// Login devuelve un JWT al usuario.
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

// Logout recibe el Refresh Token a invalidar del usuario en el body.
// Una vez invalidado el Refresh Token, el usuario no puede obtener nuevos Access Tokens.
func (h *authHandler) Logout() gin.HandlerFunc {
	type RequestBody struct {
		Token string `json:"token"`
	}

	return func(ctx *gin.Context) {
		var body RequestBody
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("failed to parse json body"))
		}
		// TODO: validar header Authorization (con middleware)
		err = h.s.LogoutUser(body.Token)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, errors.New("something went wrong"))
		}
		web.Success(ctx, http.StatusOK, nil)
	}
}

// RefreshToken recibe un Refresh Token, y devuelve un nuevo JWT.
// Esto permite al usuario conseguir un Access Token nuevo, cuando vence el previo.
func (h *authHandler) RefreshToken() gin.HandlerFunc {
	type RequestBody struct {
		GrantType    string `json:"grantType"`
		RefreshToken string `json:"refreshToken"`
	}
	return func(ctx *gin.Context) {
		var body RequestBody
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("failed to parse json body"))
		}
		// GrantType debe ser refresh_token por el estándar de OAuth.
		if body.GrantType != "refresh_token" {
			web.Failure(ctx, http.StatusBadRequest, errors.New("grantType value must be 'refresh_token'"))
		}
		// TODO: validar header Authorization (con middleware)
		jwt, err := h.s.RefreshToken(body.RefreshToken)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, errors.New("something went wrong"))
		}
		web.Success(ctx, http.StatusOK, jwt)
	}
}
