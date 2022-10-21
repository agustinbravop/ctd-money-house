package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/agustinbravop/ctd-money-house/internal/user"
	"github.com/agustinbravop/ctd-money-house/pkg/web"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	s user.Service
}

func NewUserHandler(s user.Service) *userHandler {
	return &userHandler{
		s: s,
	}
}

func (h *userHandler) GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		user, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, http.StatusNotFound, errors.New("user not found"))
			return
		}
		web.Success(c, http.StatusOK, user)
	}
}

func (h *userHandler) GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := h.s.GetAll()
		if err != nil {
			web.Failure(c, http.StatusInternalServerError, errors.New("internal server error"))
			return
		}
		web.Success(c, http.StatusOK, users)
	}
}
