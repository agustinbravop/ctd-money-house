package web

import (
	"accounts-api/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Success(ctx *gin.Context, status int, data interface{}) {
	ctx.JSON(status, data)
}

func Failure(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, errorResponse{
		Message: err.Error(),
		Status:  http.StatusText(status),
		Code:    status,
	})
}

func FailureApiErr(ctx *gin.Context, apiErr utils.ApiError) {
	ctx.JSON(apiErr.Status(), errorResponse{
		Message: apiErr.Message(),
		Status:  apiErr.Error(),
		Code:    apiErr.Status(),
	})
}

// Abort llama a Failure y luego hace ctx.Abort(). Útil para los middlewares.
func Abort(ctx *gin.Context, status int, err error) {
	Failure(ctx, status, err)
	ctx.Abort()
}
