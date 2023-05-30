package utils

import (
	"fmt"
	"net/http"
)

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiErr struct {
	ErrCode    int    `json:"status"`
	ErrMessage string `json:"message"`
	ErrCause   error  `json:"error"`
}

func (e apiErr) Status() int {
	return e.ErrCode
}

func (e apiErr) Message() string {
	return e.ErrMessage
}

func (e apiErr) Error() string {
	return fmt.Sprintf("Error: %s", e.ErrCause)
}

func NewInternalServerError(message string, err error) ApiError {
	return apiErr{
		ErrCode:    http.StatusInternalServerError,
		ErrMessage: message,
		ErrCause:   err,
	}
}

func NewNotFoundError(message string, err error) ApiError {
	return apiErr{
		ErrCode:    http.StatusNotFound,
		ErrMessage: message,
		ErrCause:   err,
	}
}

func NewBadRequestError(message string, err error) ApiError {
	return apiErr{
		ErrCode:    http.StatusBadRequest,
		ErrMessage: message,
		ErrCause:   err,
	}
}

func NewConflictError(message string, err error) ApiError {
	return apiErr{
		ErrCode:    http.StatusConflict,
		ErrMessage: message,
		ErrCause:   err,
	}
}
