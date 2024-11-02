// pkg/errors/errors.go
package errors

import (
	"fmt"
	"net/http"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("code=%d, message=%s", e.Code, e.Message)
}

func NewAppError(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

func NewNotFoundError(message string) *AppError {
	return NewAppError(http.StatusNotFound, message)
}

func NewInternalServerError(message string) *AppError {
	return NewAppError(http.StatusInternalServerError, message)
}

func NewBadRequestError(message string) *AppError {
	return NewAppError(http.StatusBadRequest, message)
}

// Repository-specific errors
func NewDatabaseError(message string) *AppError {
	return NewAppError(http.StatusInternalServerError, message)
}

func NewDataNotFoundError(message string) *AppError {
	return NewAppError(http.StatusNotFound, message)
}
