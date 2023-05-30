package routes

import (
	"bank-api/cmd/api/handlers"
	"bank-api/cmd/api/middle"
	"bank-api/internal/card"
	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
	Start(addr string) error
}

type router struct {
	eng *gin.Engine
	rg  *gin.RouterGroup
}

func NewRouter(r *gin.Engine) Router {
	return &router{
		eng: r,
	}
}

func (r *router) Start(addr string) error {
	return r.eng.Run(addr)
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.buildBankRoutes()
}

func (r *router) setGroup() {
	r.eng.Use(middle.EnableCORS())
	r.rg = r.eng.Group("/api/v1")
}

func (r *router) buildBankRoutes() {
	repo := card.NewRepository()
	service := card.NewService(repo)
	handler := handlers.NewCardHandler(service)

	bank := r.rg.Group("/bank")
	{
		bank.POST("/extraction", handler.Extraction())
	}
}
