package routes

import (
	"ctd-money-house/internal/auth"
	"database/sql"

	"ctd-money-house/cmd/api/handlers"
	"ctd-money-house/cmd/api/middleware"
	"ctd-money-house/internal/user"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{
		r:  r,
		db: db,
	}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.buildUserRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) buildUserRoutes() {
	repo := user.NewRepository(r.db)
	service := user.NewService(repo)
	handler := handlers.NewUserHandler(service)
	users := r.rg.Group("/users")
	{
		users.GET("/:id", middleware.TokenValidation(), handler.GetUserByID())
		users.GET("/", middleware.TokenValidation(), handler.GetAllUsers())
	}
}

func (r *router) buildAuthRoutes(keycloakUrl, clientId, clientSecret, realm string) {
	keycloakClient, err := auth.NewKeycloakClient(keycloakUrl, clientId, clientSecret, realm)
	// Instanciar un KeycloakClient puede fallar si la petición del cliente a Keycloak falla.
	if err != nil {
		panic(err.Error())
	}
	service := auth.NewAuthService(keycloakClient)
	handler := handlers.NewAuthHandler(service)
	auths := r.rg.Group("/auth")
	{
		auths.POST("/login", handler.Login())
	}
}
