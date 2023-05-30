package handlers

import (
	"errors"
	"net/http"
	"users-api/internal/auth"
	"users-api/internal/domain"
	"users-api/pkg/web"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	s auth.Service
}

type RequestBodyLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestBodyLogout struct {
	RefreshToken string `json:"refreshToken"`
}

type RequestBodyRefreshToken struct {
	RefreshToken string `json:"refreshToken"`
}

func NewAuthHandler(s auth.Service) *AuthHandler {
	return &AuthHandler{
		s: s,
	}
}

// Register crea un nuevo usuario, y retorna el usuario creado.
// Register godoc
// @Summary Register
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body UserReq true "body"
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse 
// @Success 201 {object} domain.User
// @Router /api/v1/auth/register [post]
func (h *AuthHandler) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body UserReq
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, ErrBadRequest)
			return
		}
		newUser := domain.User{
			Name:      body.FirstName,
			LastName:  body.LastName,
			Dni:       body.Dni,
			Email:     body.Email,
			Telephone: body.Telephone,
		}
		newUser, err = h.s.RegisterUser(newUser, body.Password)
		if err != nil {
			respondAuthErrorFailure(ctx, err)
			return
		}
		web.Success(ctx, http.StatusCreated, newUser)
	}
}

// Login devuelve un JWT con un Access Token y un Refresh Token válidos al usuario.
// Register crea un nuevo usuario, y retorna el usuario creado.
// Login godoc
// @Summary Login
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body RequestBodyLogin true "body"
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse 
// @Success 200 {object} domain.User
// @Router /api/v1/auth/login [post]
func (h *AuthHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body RequestBodyLogin
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, ErrBadRequest)
			return
		}

		jwt, err := h.s.LoginUser(body.Email, body.Password)
		if err != nil {
			respondAuthErrorFailure(ctx, err)
			return
		}
		web.Success(ctx, http.StatusOK, jwt)
	}
}

// Logout recibe el Refresh Token a invalidar del usuario en el body.
// Una vez invalidado el Refresh Token, el usuario no puede obtener nuevos Access Tokens.
// Logout godoc
// @Summary Logout
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body RequestBodyLogout true "body"
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse 
// @Success 200 
// @Router /api/v1/auth/logout [post]
func (h *AuthHandler) Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body RequestBodyLogout
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, ErrBadRequest)
			return
		}
		err = h.s.LogoutUser(body.RefreshToken)
		if err != nil {
			respondAuthErrorFailure(ctx, err)
			return
		}
		web.Success(ctx, http.StatusOK, "")
	}
}

// RefreshToken recibe un Refresh Token, y devuelve un JWT con un nuevo Access Token y Refresh Token.
// RefreshToken godoc
// @Summary Refresh Token
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param request body RequestBodyRefreshToken true "body"
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse 
// @Success 200 "returns jwt" 
// @Router /api/v1/auth/token [post]
func (h *AuthHandler) RefreshToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body RequestBodyRefreshToken
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, ErrBadRequest)
			return
		}
		jwt, err := h.s.RefreshToken(body.RefreshToken)
		if err != nil {
			respondAuthErrorFailure(ctx, err)
			return
		}
		web.Success(ctx, http.StatusOK, jwt)
	}
}

// ValidateToken valida el Access Token del 'Authentication' Header contra Keycloak y devuelve sus claims (payload).
// ValidateToken godoc
// @Summary Validate Token
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param Authorization header string true "header"
// @Failure 400 {object} web.errorResponse
// @Failure 401 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse 
// @Success 200 
// @Router /api/v1/auth/validate [post]
func (h *AuthHandler) ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.GetHeader("Authorization")

		if token == "" {
			web.Failure(ctx, 401, errors.New("access token required"))
			return
		}
		_, claims, err := h.s.DecodeToken(token)
		if err != nil {
			web.Failure(ctx, 401, errors.New("invalid token"))
			return
		}
		web.Success(ctx, http.StatusOK, claims)
	}
}

// UpdateUser es el PATCH /users/:id, que como necesita afectar a la DB de Keycloak, debe llamar al auth.Service.
// Lo que sucede ahora es que el auth.Service llama al users.Service pero en algunas ocasiones conviene que sea al revés.
// Nota: esto sugiere combinar al users.Service y al auth.Service en uno solo, para no mezclar sus funcionalidades.
// UpdateUser godoc
// @Summary Update user
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path string true "user ID"
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse 
// @Success 200 {object} domain.User
// @Router /api/v1/users/{id} [post]
func (h *AuthHandler) UpdateUser() gin.HandlerFunc {
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
		u.ID = id
		updatedUser, err := h.s.UpdateUser(u, userRequest.Password)
		if err != nil {
			respondAuthErrorFailure(ctx, err)
			return
		}
		web.Success(ctx, 200, updatedUser)
	}
}

func respondAuthErrorFailure(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, auth.ErrMissingFields):
		web.Failure(ctx, 400, err)
	case errors.Is(err, auth.ErrInternal):
		web.Failure(ctx, 500, err)
	case errors.Is(err, auth.ErrWrongEmail):
		web.Failure(ctx, 404, err)
	case errors.Is(err, auth.ErrWrongPassword):
		web.Failure(ctx, 400, err)
	case errors.Is(err, auth.ErrEmailInUse):
		web.Failure(ctx, 400, err)
	default:
		web.Failure(ctx, 500, errors.New("something went wrong"))
	}
}
