package http

import (
	"encoding/json"
	"github.com/devisions/go-mux-jwt-gorm-starter/app"
	"log"
	"net/http"
)

type ErrorContent struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	ErrorContent ErrorContent `json:"error"`
}

func newErrorResponse(err app.Error) ErrorResponse {
	return ErrorResponse{
		ErrorContent: ErrorContent{
			Code:    err.Code,
			Message: err.Message,
		},
	}
}

// RespondWithError is responding with the provided error and statusCode.
func RespondWithError(w http.ResponseWriter, statusCode int, err app.Error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	log.Printf("Failed to encode the response: %s\n", json.NewEncoder(w).Encode(newErrorResponse(err)))
}
