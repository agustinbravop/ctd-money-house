package routes

import (
	"accounts-api/cmd/api/handlers"
	"accounts-api/cmd/api/middle"
	"accounts-api/internal/account"
	"accounts-api/internal/card"
	"accounts-api/internal/transaction"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Router interface {
	MapRoutes()
	Start(addr string) error
}

type router struct {
	eng *gin.Engine
	rg  *gin.RouterGroup
	db  *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{
		eng: r,
		db:  db,
	}
}

func (r *router) Start(addr string) error {
	return r.eng.Run(addr)
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.buildAccountRoutes()
}

func (r *router) setGroup() {
	r.eng.Use(middle.EnableCORS())
	r.rg = r.eng.Group("/api/v1")
}

func (r *router) buildAccountRoutes() {
	accountRepository := account.NewRepository(r.db)
	accountService := account.NewService(accountRepository)
	accountHandler := handlers.NewAccountHandler(accountService)

	cardRepository := card.NewRepository(r.db)
	cardService := card.NewService(cardRepository, accountService)
	cardHandler := handlers.NewCardHandler(cardService)

	trRepository := transaction.NewRepository(r.db)
	trService := transaction.NewService(trRepository, accountRepository, cardRepository)
	trHandler := handlers.NewTransactionHandler(trService)

	validation := middle.New(accountService)

	accounts := r.rg.Group("/accounts", middle.AuthRequired())
	{
		accounts.GET("/", accountHandler.GetAccountByUserID())
		accounts.POST("/", accountHandler.Create())
		accounts.GET("/byAliasOrCvu/:aliasOrCvu", accountHandler.GetAccountByAliasOrCvu())

		accountID := accounts.Group("/:id", validation.AccountUserIDMatchesAuthHeader())
		{
			accountID.GET("/", accountHandler.GetAccountByID())
			accountID.PATCH("/", accountHandler.UpdateAlias())
			accountID.GET("/activity/:transactionId", trHandler.GetTransactionByID())
			accountID.GET("/activity", accountHandler.GetLastTransactions())
			accountID.GET("/activity/filter", trHandler.FilterTransactions())
			accountID.POST("/deposit", trHandler.DepositFromCard())
			accountID.POST("/transactions", trHandler.CreateTransaction())
			accountID.GET("/cards", cardHandler.GetAllCardsByAccountID())
			accountID.GET("/cards/:idCard", cardHandler.GetCardByID())
			accountID.POST("/cards", cardHandler.Create())
			accountID.DELETE("/cards/:idCard", cardHandler.Delete())
		}
	}
}
