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
			return
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
			return
		}
		web.Success(ctx, http.StatusCreated, user)
	}
}

// Login devuelve un JWT con un Access Token y un Refresh Token válidos al usuario.
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
			return
		}

		jwt, err := h.s.LoginUser(body.Email, body.Password)
		if err != nil {
			// TODO: con (err != nil) no se puede diferenciar cual fue el error.
			// TODO: sería mejor si pudieramos dar un mensaje de error más personalizado.
			println(err.Error())
			web.Failure(ctx, http.StatusBadRequest, errors.New("wrong password or email"))
			return
		}
		web.Success(ctx, http.StatusOK, jwt)
	}
}

// Logout recibe el Refresh Token a invalidar del usuario en el body.
// Una vez invalidado el Refresh Token, el usuario no puede obtener nuevos Access Tokens.
func (h *authHandler) Logout() gin.HandlerFunc {
	type RequestBody struct {
		RefreshToken string `json:"refreshToken"`
	}

	return func(ctx *gin.Context) {
		var body RequestBody
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("failed to parse json body"))
			return
		}
		err = h.s.LogoutUser(body.RefreshToken)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, errors.New("something went wrong"))
			return
		}
		web.Success(ctx, http.StatusOK, nil)
	}
}

// RefreshToken recibe un Refresh Token, y devuelve un JWT con un nuevo Access Token y Refresh Token.
func (h *authHandler) RefreshToken() gin.HandlerFunc {
	type RequestBody struct {
		RefreshToken string `json:"refreshToken"`
	}

	return func(ctx *gin.Context) {
		var body RequestBody
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("failed to parse json body"))
			return
		}
		jwt, err := h.s.RefreshToken(body.RefreshToken)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, errors.New("something went wrong"))
			return
		}
		web.Success(ctx, http.StatusOK, jwt)
	}
}
