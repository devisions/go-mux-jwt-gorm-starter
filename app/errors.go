package app

import "errors"

type Error struct {
	Code    string
	Message string
}

type ErrorType string

func (et ErrorType) AsError() error {
	return errors.New(string(et))
}

func GetError(errorType ErrorType) Error {
	return appErrors[errorType]
}

var appErrors map[ErrorType]Error = map[ErrorType]Error{
	ErrEmpty:         {},
	ErrInternal:      {"E001", "An internal error occurred."},
	ErrUnauthorized:  {"E002", "Unauthorized access."},
	ErrJWTValidation: {"E003", "JWT validation error."},
	ErrNotFound:      {"E004", "Not found."},
	ErrInvalidCreds:  {"E005", "Invalid credentials."},
	ErrBadRequest:    {"E006", "Invalid request."},
}

const (
	ErrEmpty         ErrorType = "ErrEmpty"
	ErrInternal      ErrorType = "ErrInternal"
	ErrUnauthorized  ErrorType = "ErrUnauthorized"
	ErrJWTValidation ErrorType = "ErrJWTValidation"
	ErrNotFound      ErrorType = "ErrNotFound"
	ErrInvalidCreds  ErrorType = "ErrInvalidCreds"
	ErrBadRequest    ErrorType = "ErrBadRequest"
)
