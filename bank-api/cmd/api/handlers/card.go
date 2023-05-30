package handlers

import (
	"bank-api/internal/card"
	"bank-api/internal/domain"
	"bank-api/pkg/web"
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrBadRequest = errors.New("invalid json body")
)

type ExtractionReq struct {
	CardNumber     string  `json:"card_number"`
	ExpiryDate     string  `json:"expiry_date"`
	Owner          string  `json:"owner"`
	SecurityCode   string  `json:"security_code"`
	Amount         float64 `json:"amount"`
	DestinationCvu string  `json:"destination_cvu"`
}

type ExtractionRes struct {
	Amount         float64 `json:"amount"`
	OriginCvu      string  `json:"origin_cvu"`
	DestinationCvu string  `json:"destination_cvu"`
}

type CardHandler struct {
	s card.Service
}

func NewCardHandler(s card.Service) *CardHandler {
	return &CardHandler{
		s: s,
	}
}

func (h *CardHandler) Extraction() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var extractionReq ExtractionReq
		if err := ctx.ShouldBindJSON(&extractionReq); err != nil {
			web.Failure(ctx, 400, ErrBadRequest)
			return
		}

		c := createDomainCard(&extractionReq)
		extraction, err := h.s.Extract(extractionReq.Amount, extractionReq.DestinationCvu, c)
		if err != nil {
			respondCardErrorFailure(ctx, err)
			return
		}
		web.Success(ctx, 200, extraction)
	}
}

func createDomainCard(card *ExtractionReq) domain.Card {
	return domain.Card{
		CardNumber:   card.CardNumber,
		SecurityCode: card.SecurityCode,
		Owner:        card.Owner,
		ExpiryDate:   card.ExpiryDate,
	}
}

// respondCardErrorFailure matchea contra todos los tipos de error definidos en el paquete card, y responde acorde.
func respondCardErrorFailure(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, card.ErrInternal):
		web.Failure(ctx, 500, err)
	case errors.Is(err, card.ErrInsufficientFunds):
		web.Failure(ctx, 409, err)
	case errors.Is(err, card.ErrExpiredCard):
		web.Failure(ctx, 409, err)
	case errors.Is(err, card.ErrBlockedCard):
		web.Failure(ctx, 409, err)
	case errors.Is(err, card.ErrNotFound):
		web.Failure(ctx, 404, err)
	case errors.Is(err, card.ErrDifferentFields):
		web.Failure(ctx, 403, err)
	default:
		web.Failure(ctx, 500, errors.New("something went wrong"))
	}
}
