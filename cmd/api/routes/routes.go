package routes

import (
	"database/sql"

	"github.com/agustinbravop/ctd-money-house/cmd/api/handlers"
	"github.com/agustinbravop/ctd-money-house/cmd/api/middleware"
	"github.com/agustinbravop/ctd-money-house/internal/user"
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
