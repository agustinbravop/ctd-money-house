package handlers

import (
	"errors"
	"net/http"
	"users-api/internal/domain"
	"users-api/internal/user"
	"users-api/pkg/web"

	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidID    = errors.New("invalid id")
	ErrMismatchedID = errors.New("user id in token is different to user id in url")
	ErrBadRequest   = errors.New("invalid json body")
)

type UserReq struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Dni       string `json:"dni"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
}

type UserHandler struct {
	s user.Service
}

func NewUserHandler(s user.Service) *UserHandler {
	return &UserHandler{
		s: s,
	}
}

// GetUserByID godoc
// @Summary Get user by ID
// @Tags User
// @Accept  json
// @Produce  json
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse 
// @Success 200 {array} domain.User
// @Router /api/v1/users [get]
func (h *UserHandler) GetUserByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !isUrlIDValid(ctx) {
			return
		}
		id := ctx.Param("id")
		foundUser, err := h.s.GetByID(id)
		if err != nil {
			respondUserErrorFailure(ctx, err)
			return
		}

		web.Success(ctx, http.StatusOK, foundUser)
	}
}

// GetAllUsers godoc
// @Summary Get all users
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "user ID"
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse 
// @Success 200 {object} domain.User
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := h.s.GetAll()
		if err != nil {
			respondUserErrorFailure(c, err)
			return
		}
		web.Success(c, http.StatusOK, users)
	}
}

func (h *UserHandler) UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userRequest := UserReq{}
		if err := ctx.ShouldBind(&userRequest); err != nil {
			web.Failure(ctx, 400, ErrBadRequest)
			return
		}
		id := ctx.Param("id")
		if !isUrlIDValid(ctx) {
			return
		}
		u := createDomainUser(userRequest)

		// ctx.GetString() obtiene el valor de UserID, seteado por el middle previo AuthRequired.
		userResponse, err := h.s.Update(id, u)
		if err != nil {
			respondUserErrorFailure(ctx, err)
			return
		}
		web.Success(ctx, 200, userResponse)
	}
}

// Create godoc
// @Summary Create user
// @Tags User
// @Accept  json
// @Produce  json
// @Param request body UserReq true "body"
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse 
// @Success 201 {object} domain.User
// @Router /api/v1/users [post]
func (h *UserHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userReq UserReq

		if err := ctx.ShouldBindJSON(&userReq); err != nil {
			web.Failure(ctx, 400, ErrBadRequest)
			return
		}
		u := createDomainUser(userReq)

		createdUser, err := h.s.Create(u)
		if err != nil {
			respondUserErrorFailure(ctx, err)
			return
		}

		web.Success(ctx, 201, createdUser)
	}
}

// Delete godoc
// @Summary Delete user
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "user ID"
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse 
// @Success 204
// @Router /api/v1/users/{id} [delete]
func (h *UserHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !isUrlIDValid(ctx) {
			return
		}

		id := ctx.Param("id")
		err := h.s.Delete(id)
		if err != nil {
			respondUserErrorFailure(ctx, err)
			return
		}
		web.Success(ctx, 204, "")
	}
}

func createDomainUser(user UserReq) domain.User {
	return domain.User{
		Name:      user.FirstName,
		LastName:  user.LastName,
		Dni:       user.Dni,
		Email:     user.Email,
		Telephone: user.Telephone,
	}
}

// isUrlIDValid valida que el userID de la URL sea igual al UserID del token, y devuelve true si coinciden.
// El UserID del Token fue guardado en la Key 'UserID' del gin.Context por el middleware AuthRequired.
// Si no coinciden, responde con un 400 (usa web.Failure) y devuelve false.
func isUrlIDValid(ctx *gin.Context) bool {
	id := ctx.Param("id")
	if id == "" {
		web.Failure(ctx, 400, ErrInvalidID)
		return false
	}
	if id != ctx.GetString("UserID") {
		web.Failure(ctx, 400, ErrMismatchedID)
		return false
	}
	return true
}

// respondUserErrorFailure matchea contra todos los tipos de error definidos en el paquete user, y responde acorde.
func respondUserErrorFailure(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, user.ErrNotFound):
		web.Failure(ctx, 404, err)
	case errors.Is(err, user.ErrMissingFields):
		web.Failure(ctx, 400, err)
	case errors.Is(err, user.ErrInternal):
		web.Failure(ctx, 500, err)
	case errors.Is(err, user.ErrEmailInUse):
		web.Failure(ctx, 400, err)
	default:
		web.Failure(ctx, 500, errors.New("something went wrong"))
	}
}
