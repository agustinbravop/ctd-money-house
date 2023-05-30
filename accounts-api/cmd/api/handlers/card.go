package handlers

import (
	"net/http"
	"strconv"

	"accounts-api/internal/card"
	"accounts-api/internal/domain"
	"accounts-api/pkg/utils"
	"accounts-api/pkg/web"

	"github.com/gin-gonic/gin"
)

type cardHandler struct {
	s card.Service
}

func NewCardHandler(s card.Service) *cardHandler {
	return &cardHandler{
		s: s,
	}
}

// GetCardByID godoc
// @Summary Get card by ID
// @Tags Card
// @Accept  json
// @Produce  json
// @Param idCard path int true "card ID"
// @Param id path int true "account ID"
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse 
// @Success 200 {object} domain.Card
// @Router /api/v1/accounts/{id}/cards/{idCard} [get]
func (h *cardHandler) GetCardByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("idCard"))
		if err != nil {
			web.Failure(c, http.StatusBadRequest, ErrInvalidID)
			return
		}
		card, apiErr := h.s.GetByID(id)
		if apiErr != nil {
			web.FailureApiErr(c, apiErr)
			return
		}

		web.Success(c, http.StatusOK, card)
	}
}


// GetAllCardsByAccountID godoc
// @Summary Get all cards by account ID
// @Tags Card
// @Accept  json
// @Produce  json
// @Param id path int true "account ID"
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse 
// @Success 200 {array} domain.Card
// @Router /api/v1/accounts/{id}/cards [get]
func (h *cardHandler) GetAllCardsByAccountID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, http.StatusBadRequest, ErrInvalidID)
			return
		}
		cards, apiErr := h.s.GetAllByAccountID(id)
		if apiErr != nil {
			web.FailureApiErr(c, apiErr)
			return
		}
		web.Success(c, http.StatusOK, cards)
	}
}

// Create godoc
// @Summary Create card
// @Tags Card
// @Accept  json
// @Produce  json
// @Param id path int true "account ID"
// @Param request body domain.Card true "body"
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse 
// @Success 201 {object} domain.Card
// @Router /api/v1/accounts/{id}/cards [post]
func (h *cardHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cardReq domain.Card

		accountID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, http.StatusBadRequest, ErrInvalidID)
			return
		}

		cardReq.AccountID = accountID
		if err := c.ShouldBindJSON(&cardReq); err != nil {
			web.Failure(c, 400, ErrBadRequest)
			return
		}

		if utils.AnyBlank(cardReq.CardNumber, cardReq.Owner, cardReq.ExpirationDate, cardReq.Brand, cardReq.SecurityCode) {
			web.Failure(c, http.StatusBadRequest, ErrBadRequest)
			return
		}

		card, apiErr := h.s.Create(cardReq)
		if apiErr != nil {
			web.FailureApiErr(c, apiErr)
			return
		}

		web.Success(c, 201, card)
	}
}


// Delete godoc
// @Summary Delete card
// @Tags Card
// @Accept  json
// @Produce  json
// @Param id path int true "account ID"
// @Failure 400 {object} web.errorResponse
// @Failure 500 {object} web.errorResponse 
// @Success 204 
// @Router /api/v1/accounts/{id}/cards [delete]
func (h *cardHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

		accountID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, 400, ErrInvalidID)
			return
		}

		cardID, err := strconv.Atoi(c.Param("idCard"))
		if err != nil {
			web.Failure(c, 400, ErrInvalidID)
			return
		}

		if apiErr := h.s.Delete(int(accountID), int(cardID)); apiErr != nil {
			web.FailureApiErr(c, apiErr)
			return
		}
		web.Success(c, 204, "")
	}
}
