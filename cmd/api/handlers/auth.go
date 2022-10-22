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

func (h *authHandler) Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var jwt auth.JWT
		err := ctx.ShouldBindJSON(&jwt)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("failed to parse json body"))
		}

		err = h.s.LogoutUser(jwt)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, errors.New("something went wrong"))
		}
		web.Success(ctx, http.StatusOK, nil)
	}
}
