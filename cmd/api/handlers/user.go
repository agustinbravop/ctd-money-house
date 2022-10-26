package handlers

import (
	"ctd-money-house/internal/domain"
	"ctd-money-house/internal/user"
	"ctd-money-house/pkg/web"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidID    = errors.New("invalid id")
	ErrBadRequest   = errors.New("invalid json body")
	ErrUserNotFound = errors.New("user not found")
)

type UserReq struct {
	Name      string `json:"first_name"`
	LastName  string `json:"last_name"`
	Dni       string `json:"dni"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
}

type UserRes struct {
	Name      string `json:"first_name"`
	LastName  string `json:"last_name"`
	Dni       string `json:"dni"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Cvu       string `json:"cvu"`
	Alias     string `json:"alias"`
}

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
			web.Failure(c, http.StatusBadRequest, ErrInvalidID)
			return
		}
		user, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, http.StatusNotFound, ErrUserNotFound)
			return
		}

		userResp := UserRes{
			Name:      user.Name,
			LastName:  user.LastName,
			Dni:       user.Dni,
			Email:     user.Email,
			Telephone: user.Telephone,
			Cvu:       user.Cvu,
			Alias:     user.Alias,
		}

		web.Success(c, http.StatusOK, userResp)
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

func (h *userHandler) UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRequest := userReq{}
		if err := c.ShouldBind(&userRequest); err != nil {
			web.Failure(c, 400, errors.New("error should bind"))
			return
		}

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}

		u := createDomainUser(userRequest)

		userResponse, err := h.s.Update(int(id), u)
		switch err {
		case nil:
			web.Success(c, 200, userResponse)
			return
		case user.ErrInternal:
			web.Failure(c, 500, err)
			return
		case user.ErrNotFound:
			web.Failure(c, 404, err)
			return
		default:
			web.Failure(c, 404, errors.New("try again later"))
			return
		}
	}
}
func (h *userHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userReq UserReq

		if err := c.ShouldBindJSON(&userReq); err != nil {
			web.Failure(c, 400, ErrBadRequest)
			return
		}

		u := createDomainUser(userReq)

		resp, err := h.s.Create(u)
		if err != nil {
			switch {
			case errors.Is(err, user.ErrInternal):
				web.Failure(c, 500, err)
				return
			default:
				web.Failure(c, 409, err)
				return
			}
		}

		res := UserRes{
			Name:      resp.Name,
			LastName:  resp.LastName,
			Dni:       resp.Dni,
			Email:     resp.Email,
			Telephone: resp.Telephone,
			Alias:     resp.Alias,
			Cvu:       resp.Cvu,
		}

		web.Success(c, 201, res)
	}
}

func (h *userHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 32)
		if err != nil {
			web.Failure(c, 400, ErrInvalidID)
			return
		}
		err = h.s.Delete(int(id))
		if err != nil {
			switch {
			case errors.Is(err, user.ErrInternal):
				web.Failure(c, 500, err)
				return
			default:
				web.Failure(c, 404, err)
				return
			}
		}
		web.Success(c, 204, "")
	}
}

func createDomainUser(user userReq) domain.User {
	return domain.User{
		Name: user.Name,
		LastName: user.LastName,
		Dni: user.Dni,
		Email: user.Email,
		Telephone: user.Telephone,
	}
}