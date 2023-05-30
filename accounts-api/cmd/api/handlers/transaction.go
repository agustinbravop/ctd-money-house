package handlers

import (
	"accounts-api/internal/domain"
	"accounts-api/internal/transaction"
	"accounts-api/pkg/web"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidFilter = errors.New("invalid filter. Only allowed filters: amount, date, type")
)

type TransactionReq struct {
	Amount          float64 `json:"amount"`
	Description     string  `json:"description"`
	OriginCvu       string  `json:"origin_cvu"`
	DestinationCvu  string  `json:"destination_cvu"`
	AccountID       int
	TransactionType string `json:"transaction_type"`
}

type DepositReq struct {
	Amount float64 `json:"amount"`
	CardID int     `json:"card_id"`
}

type transactionHandler struct {
	s transaction.Service
}

func NewTransactionHandler(s transaction.Service) *transactionHandler {
	return &transactionHandler{
		s: s,
	}
}

// GetTransactionByID godoc
// @Summary Get transaction by ID
// @Tags Transaction
// @Accept  json
// @Produce  json
// @Param id path int true "account ID"
// @Param transactionId path int true "transaction ID"
// @Failure 400 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Failure 403 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Success 200 {object} domain.Transaction
// @Router /api/v1/accounts/{id}/activity/{transactionId} [get]
func (h *transactionHandler) GetTransactionByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		transactionId, err := strconv.Atoi(c.Param("transactionId"))
		if err != nil {
			web.Failure(c, http.StatusBadRequest, ErrInvalidID)
			return
		}
		foundTransaction, err := h.s.GetByID(transactionId)
		if err != nil {
			respondTransactionErrorFailure(c, err)
			return
		}

		web.Success(c, http.StatusOK, foundTransaction)
	}
}

// GetLastTransactions godoc
// @Summary Get last transaction
// @Tags Transaction
// @Accept  json
// @Produce  json
// @Param id path int true "account ID"
// @Param id query int true "limit"
// @Failure 400 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Failure 403 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Success 200 {array} domain.Transaction
// @Router /api/v1/accounts/{id}/activity [get]
func (h *transactionHandler) GetLastTransactions() gin.HandlerFunc {
	return func(c *gin.Context) {
		accountId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, http.StatusBadRequest, ErrInvalidID)
			return
		}
		transactions, err := h.s.GetAllByAccountID(accountId)
		if err != nil {
			respondTransactionErrorFailure(c, err)
			return
		}
		web.Success(c, http.StatusOK, transactions)
	}
}

// FilterTransactions godoc
// @Summary Filter transactions
// @Tags Transaction
// @Produce  json
// @Param id path int true "account ID"
// @Param id query string true "type"
// @Param id query string false "from"
// @Param id query string false "to"
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Success 200 {array} domain.Transaction
// @Router /api/v1/accounts/{id}/activity/filter [get]
func (h *transactionHandler) FilterTransactions() gin.HandlerFunc {
	return func(c *gin.Context) {
		accountID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, http.StatusBadRequest, ErrInvalidID)
			return
		}

		var filter domain.Filter
		if err = c.ShouldBindQuery(&filter); err != nil {
			web.Failure(c, http.StatusBadRequest, ErrInvalidFilter)
			return
		}
		log.Printf("%#+v\n", filter)

		transactions, apiErr := h.s.FilterTransactions(accountID, filter)
		if apiErr != nil {
			web.FailureApiErr(c, apiErr)
			return
		}
		web.Success(c, http.StatusOK, transactions)
	}
}

// CreateTransaction godoc
// @Summary Create transaction
// @Tags Transaction
// @Accept  json
// @Produce  json
// @Param id path int true "account ID"
// @Param request body TransactionReq true "body - cvu and alias"
// @Failure 400 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Failure 403 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Success 201 {object} domain.Transaction
// @Router /api/v1/accounts/{id}/transactions [post]
func (h *transactionHandler) CreateTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var transactionReq TransactionReq
		accountID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, http.StatusBadRequest, ErrInvalidID)
			return
		}
		if err := c.ShouldBindJSON(&transactionReq); err != nil {
			web.Failure(c, 400, ErrBadRequest)
			return
		}
		transactionReq.AccountID = accountID
		tr := createInternalTransaction(transactionReq)
		createdTransaction, err := h.s.Create(tr)
		if err != nil {
			respondTransactionErrorFailure(c, err)
			return
		}

		web.Success(c, 201, createdTransaction)
	}
}

// DepositFromCard godoc
// @Summary Deposit from card
// @Tags Transaction
// @Accept  json
// @Produce  json
// @Param id path int true "account ID"
// @Param request body TransactionReq true "body"
// @Failure 400 {object} web.errorResponse
// @Failure 409 {object} web.errorResponse
// @Failure 403 {object} web.errorResponse
// @Failure 404 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse
// @Success 200 {array} domain.Transaction
// @Router /api/v1/accounts/{id}/deposit [post]
func (h *transactionHandler) DepositFromCard() gin.HandlerFunc {
	return func(c *gin.Context) {
		accountID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, http.StatusBadRequest, ErrInvalidID)
			return
		}

		var depositReq DepositReq
		if err := c.ShouldBindJSON(&depositReq); err != nil {
			web.Failure(c, 400, ErrBadRequest)
			return
		}

		createdTransaction, err := h.s.DepositFromCard(accountID, depositReq.Amount, depositReq.CardID)
		if err != nil {
			respondTransactionErrorFailure(c, err)
			return
		}

		web.Success(c, 201, createdTransaction)
	}
}

func createInternalTransaction(t TransactionReq) domain.Transaction {
	return domain.Transaction{
		Amount:          t.Amount,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
		Description:     t.Description,
		OriginCvu:       t.OriginCvu,
		DestinationCvu:  t.DestinationCvu,
		AccountID:       t.AccountID,
		TransactionType: t.TransactionType,
	}
}

func respondTransactionErrorFailure(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, transaction.ErrInternal):
		web.Failure(ctx, 500, err)
	case errors.Is(err, transaction.ErrGettingTransaction):
		web.Failure(ctx, 500, err)
	case errors.Is(err, transaction.ErrExtractionFromCard):
		web.Failure(ctx, 500, err)
	case errors.Is(err, transaction.ErrInsufficientFunds):
		web.Failure(ctx, 409, err)
	case errors.Is(err, transaction.ErrNotFound):
		web.Failure(ctx, 404, err)
	case errors.Is(err, transaction.ErrAccountNotFound):
		web.Failure(ctx, 404, err)
	case errors.Is(err, transaction.ErrDestinationAccountNotFound):
		web.Failure(ctx, 404, err)
	case errors.Is(err, transaction.ErrCardNotFound):
		web.Failure(ctx, 404, err)
	case errors.Is(err, transaction.ErrMismatchedAccountID):
		web.Failure(ctx, 403, err)
	case errors.Is(err, transaction.ErrAmountCannotBeNegative):
		web.Failure(ctx, 400, err)
	case errors.Is(err, transaction.ErrSameAccount):
		web.Failure(ctx, 400, err)
	case errors.Is(err, transaction.ErrMisingFields):
		web.Failure(ctx, 400, err)
	default:
		web.Failure(ctx, 500, errors.New("something went wrong"))
	}
}
