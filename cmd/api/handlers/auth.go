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
			web.Failure(ctx, http.StatusBadRequest, errors.New("bad request"))
		}

	}
}
