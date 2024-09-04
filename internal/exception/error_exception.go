package exception

import (
	"errors"
)

var (
	ErrNotFound             = errors.New("record not found")
	ErrDuplicateEntry       = errors.New("duplicate entry")
	ErrInvalidData          = errors.New("invalid data")
	ErrInternal             = errors.New("internal server error")
	ErrUnauthorized         = errors.New("unauthorized access")
	ErrInsufficientQuantity = errors.New("insufficient quantity")
	ErrInvalidToken         = errors.New("invalid token")
	ErrTokenNotFound        = errors.New("token not found")
	ErrSigningMethodFailed  = errors.New("signing method failed")
	ErrParseClaimed         = errors.New("failed to parse claim")
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrTokenExpired         = errors.New("invalid or expired refresh token")
)

type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"err"`
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NewAppError(code, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
