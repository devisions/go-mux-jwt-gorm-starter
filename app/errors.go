package app

type Error struct {
	Code    string
	Message string
}

type ErrorType string

type Errors struct {
	catalog map[ErrorType]Error
}

// AppErrors is a catalog of application wide errors.
var errors Errors = InitErrors()

func GetError(errorType ErrorType) Error {
	return errors.catalog[errorType]
}

func InitErrors() Errors {
	catalog := map[ErrorType]Error{
		InternalError:      {"E001", "An internal error occurred."},
		UnauthorizedError:  {"E002", "Unauthorized access."},
		JWTValidationError: {"E003", "JWT validation error."},
	}
	return Errors{catalog}
}

const (
	InternalError      ErrorType = "InternalError"
	UnauthorizedError  ErrorType = "UnauthorizedError"
	JWTValidationError ErrorType = "JWTValidationError"
)
