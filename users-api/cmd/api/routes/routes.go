package routes

import (
	"database/sql"
	"users-api/internal/auth"

	"users-api/cmd/api/handlers"
	"users-api/cmd/api/middle"
	"users-api/internal/user"

	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
	Start()
}

type router struct {
	eng      *gin.Engine
	rg       *gin.RouterGroup
	db       *sql.DB
	kcClient auth.KeycloakClient
}

func NewRouter(r *gin.Engine, db *sql.DB, kcClient auth.KeycloakClient) Router {
	return &router{
		eng:      r,
		db:       db,
		kcClient: kcClient,
	}
}

func (r *router) Start() {
	err := r.eng.Run(":8082")
	if err != nil {
		panic(err)
	}
}

func (r *router) MapRoutes() {
	r.setGroup()
	userRepo := user.NewRepository(r.db)
	userService := user.NewService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	authService := auth.NewAuthService(r.kcClient, userService)
	authRequiredMiddle := middle.New(authService)
	authHandler := handlers.NewAuthHandler(authService)
	r.buildUserRoutes(userHandler, authHandler, authRequiredMiddle)
	r.buildAuthRoutes(authHandler)
}

func (r *router) setGroup() {
	r.eng.Use(middle.EnableCORS())
	r.rg = r.eng.Group("/api/v1")
}

func (r *router) buildUserRoutes(handler *handlers.UserHandler, authHandler *handlers.AuthHandler, middle *middle.AuthRequired) {
	users := r.rg.Group("/users", middle.AuthRequired())
	{
		users.GET("/:id", handler.GetUserByID())
		users.GET("/", handler.GetAllUsers())
		// PATCH /users/id llama al método UpdateUser del AuthHandler, no del UserHandler, pq puede cambiar la contraseña (Keycloak).
		users.PATCH("/:id", authHandler.UpdateUser())
		users.POST("/", handler.Create())
		users.DELETE("/:id", handler.Delete())
	}
}

func (r *router) buildAuthRoutes(handler *handlers.AuthHandler) {
	auths := r.rg.Group("/auth")
	{
		auths.POST("/register", handler.Register())
		auths.POST("/login", handler.Login())
		auths.POST("/logout", handler.Logout())
		auths.POST("/token", handler.RefreshToken())
		auths.POST("/validate", handler.ValidateToken())
	}
}
