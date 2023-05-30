package handlers

import (
	"accounts-api/internal/account"
	"accounts-api/internal/domain"
	"accounts-api/pkg/web"
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidID             = errors.New("invalid id in url")
	ErrBadRequest            = errors.New("invalid json body")
	ErrInvalidAlias          = errors.New("invalid alias")
	ErrAccountNotFound       = errors.New("account not found")
	ErrAccountNoTransactions = errors.New("this account doesn't have any transactions")
)

type AccountReq struct {
	Cvu   string `json:"cvu"`
	Alias string `json:"alias"`
}

type aliasReq struct {
	Alias *string `json:"alias"`
}

type accountHandler struct {
	s account.Service
}

func NewAccountHandler(s account.Service) *accountHandler {
	return &accountHandler{
		s: s,
	}
}

// GetAccountByID godoc
// @Summary Get account by ID
// @Tags Account
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the account to get"
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse 
// @Success 200 {object} domain.Account
// @Router /api/v1/accounts/{id} [get]
func (h *accountHandler) GetAccountByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, http.StatusBadRequest, ErrInvalidID)
			return
		}

		acc, err := h.s.GetByID(id)
		if err != nil {
			respondAcountErrorFailure(c, err)
			return
		}
		web.Success(c, http.StatusOK, acc)
	}
}

// GetAccountByAliasOrCvu godoc
// @Summary Get account by alias or cvu
// @Tags Account
// @Accept  json
// @Produce  json
// @Param aliasOrCvu path string true "alias or cvu"
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse 
// @Success 200 {object} domain.Account
// @Router /api/v1/accounts/byAliasOrCvu/{aliasOrCvu} [get]
func (h *accountHandler) GetAccountByAliasOrCvu() gin.HandlerFunc {
	return func(c *gin.Context) {
		aliasOrCvu := c.Param("aliasOrCvu")
		acc, err := h.s.GetByAliasOrCvu(aliasOrCvu)
		if err != nil {
			respondAcountErrorFailure(c, err)
			return
		}
		response := AccountReq{
			Cvu:   acc.Cvu,
			Alias: acc.Alias,
		}
		web.Success(c, http.StatusOK, response)
	}
}

// GetAccountByUserID godoc
// @Summary Get account by user ID
// @Tags Account
// @Accept  json
// @Produce  json
// @Param UserID path string true "user ID"
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse 
// @Success 200 {object} domain.Account
// @Router /api/v1/accounts/{UserID} [get]
func (h *accountHandler) GetAccountByUserID() gin.HandlerFunc {
	return func(c *gin.Context) {

		userID := c.GetString("UserID")
		acc, err := h.s.GetByUserID(userID)
		if err != nil {
			respondAcountErrorFailure(c, err)
			return
		}
		web.Success(c, http.StatusOK, acc)
	}
}

// GetLastTransactions godoc
// @Summary Get last transactions
// @Tags Account
// @Accept  json
// @Produce  json
// @Param id path int true "account ID"
// @Failure 400 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse 
// @Success 200 {array} domain.Transaction
// @Router /api/v1/accounts/activity/{id} [get]
func (h *accountHandler) GetLastTransactions() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, http.StatusBadRequest, ErrInvalidID)
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, http.StatusNotFound, ErrAccountNotFound)
			return
		}

		limit := GetLimit(c)
		transactions, err := h.s.GetLastTransactions(id, limit)
		if err != nil {
			respondAcountErrorFailure(c, err)
			return
		}
		if len(transactions) == 0 {
			web.Failure(c, http.StatusNotFound, ErrAccountNoTransactions)
			return
		}
		web.Success(c, http.StatusOK, transactions)
	}
}

// Create godoc
// @Summary Create account
// @Tags Account
// @Accept  json
// @Produce  json
// @Param request body AccountReq true "body - cvu and alias"
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse 
// @Success 201 {object} domain.Account
// @Router /api/v1/accounts [post]
func (h *accountHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var accountReq AccountReq
		if err := c.ShouldBindJSON(&accountReq); err != nil {
			web.Failure(c, 400, ErrBadRequest)
			return
		}

		acc := domain.Account{
			UserID: c.GetString("UserID"),
		}

		createdAccount, err := h.s.Create(acc)
		if err != nil {
			respondAcountErrorFailure(c, err)
			return
		}
		web.Success(c, 201, createdAccount)
	}
}

// Update godoc
// @Summary Update account
// @Tags Account
// @Accept  json
// @Produce  json
// @Param request body aliasReq true "body - alias"
// @Param id path string true "account ID"
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse 
// @Success 200 {array} domain.Account
// @Router /api/v1/accounts/{id} [patch]
func (h *accountHandler) UpdateAlias() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, 400, ErrInvalidID)
			return
		}
		var alias aliasReq
		if err := c.ShouldBindJSON(&alias); err != nil {
			web.Failure(c, 400, ErrBadRequest)
			return
		}

		if match, err := regexp.MatchString("^[a-zA-Z0-9.]+$", *alias.Alias); err != nil || !match {
			web.Failure(c, http.StatusBadRequest, ErrInvalidAlias)
			return
		}

		updatedAccount, err := h.s.UpdateAlias(id, strings.ToUpper(*alias.Alias))
		if err != nil {
			respondAcountErrorFailure(c, err)
			return
		}
		web.Success(c, 200, updatedAccount)
	}
}

func respondAcountErrorFailure(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, account.ErrInternal):
		web.Failure(ctx, 500, err)
	case errors.Is(err, account.ErrGettingAccount):
		web.Failure(ctx, 500, err)
	case errors.Is(err, account.ErrAliasConflict):
		web.Failure(ctx, 409, err)
	case errors.Is(err, account.ErrNotFound):
		web.Failure(ctx, 404, err)
	case errors.Is(err, account.ErrMissingFields):
		web.Failure(ctx, 400, err)
	default:
		web.Failure(ctx, 500, errors.New("something went wrong"))
	}
}

// GetLimit devuelve el numero del queryParam 'limit'. Si no existe, devuelve 0.
func GetLimit(c *gin.Context) uint {
	param, ok := c.GetQuery("limit")
	if !ok || strings.EqualFold(param, "") {
		return 0
	}
	limit, err := strconv.Atoi(param)
	if err != nil {
		return 0
	}
	return uint(limit)
}
