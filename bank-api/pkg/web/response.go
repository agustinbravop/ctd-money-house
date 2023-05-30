package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Success se usa cuando la response es exitosa, y solo devuelve la data dada.
func Success(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, data)
}

// Failure se usa para el body de una response fallada, y devuelve el error dado de manera estructurada.
func Failure(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, errorResponse{
		Message: err.Error(),
		Status:  status,
		Code:    http.StatusText(status),
	})
}
