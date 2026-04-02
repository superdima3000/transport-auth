package service

import "net/http"

type AppError struct {
	Err        error
	Message    string
	StatusCode int
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NewAppError(err error, message string, statusCode int) *AppError {
	return &AppError{Err: err, Message: message, StatusCode: statusCode}
}

// Sentinel errors
var (
	ErrKeyNotFound         = NewAppError(nil, "key not found", http.StatusBadRequest)
	ErrCardNotFound        = NewAppError(nil, "card not found", http.StatusBadRequest)
	ErrTerminalNotFound    = NewAppError(nil, "terminal not found", http.StatusBadRequest)
	ErrValidationFailed    = NewAppError(nil, "validation failed", http.StatusBadRequest)
	ErrCardNumberInvalid   = NewAppError(nil, "card number is invalid", http.StatusBadRequest)
	ErrNotEnoughFund       = NewAppError(nil, "not enough fund", http.StatusBadRequest)
	ErrTransactionNotFound = NewAppError(nil, "transaction not found", http.StatusBadRequest)
	ErrInvalidAmount       = NewAppError(nil, "invalid amount", http.StatusBadRequest)
	ErrUserNotFound        = NewAppError(nil, "user not found", http.StatusBadRequest)
)
