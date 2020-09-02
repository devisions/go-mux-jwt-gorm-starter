package responses

import (
	"encoding/json"
	"github.com/devisions/go-mux-jwt-gorm-starter/app"
	"log"
	"net/http"
)

type ErrorContent struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
}

type ErrorResponse struct {
	ErrorContent ErrorContent `json:"error"`
}

func newErrorResponse(errorType app.ErrorType, reason string) ErrorResponse {
	err := app.GetError(errorType)
	return ErrorResponse{
		ErrorContent{
			Code:    err.Code,
			Message: err.Message,
			Reason:  reason,
		},
	}
}

// RespondWithError responds with the provided statusCode and error in the body.
func RespondWithError(w http.ResponseWriter, statusCode int, errorType app.ErrorType) {
	RespondWithErrorReason(w, statusCode, errorType, "")
}

// RespondWithErrorReason responds with the provided statusCode and error in the body,
// including the reason, as further details.
func RespondWithErrorReason(w http.ResponseWriter, statusCode int, errorType app.ErrorType, reason string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(newErrorResponse(errorType, reason))
	if err != nil {
		log.Printf("Failed to encode the response: %s\n", err)
	}
}

// RespondAsJson responds with OK status code and JSON as body and content type.
func RespondAsJson(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(body)
}
