package responses

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/devisions/go-mux-jwt-gorm-starter/app"
)

type ErrorContent struct {
	Code    string `json:"code"`
	Message string `json:"message,omitempty"`
	Reason  string `json:"reason,omitempty"`
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

// === Common responses ===

func RespondJsonWithInternalServerError(w http.ResponseWriter) {
	RespondJsonWithErrorReason(w, http.StatusInternalServerError, app.ErrInternal, "")
}

func RespondJsonWithUnauthorizedError(w http.ResponseWriter) {
	RespondJsonWithErrorReason(w, http.StatusUnauthorized, app.ErrUnauthorized, "")
}

// === Generic responses ===

// RespondJsonWithError responds with the provided statusCode and error in the body.
func RespondJsonWithError(w http.ResponseWriter, statusCode int, errorType app.ErrorType) {
	RespondJsonWithErrorReason(w, statusCode, errorType, "")
}

// RespondJsonWithErrorReason responds with the provided statusCode and error in the body,
// including the reason, as further details.
func RespondJsonWithErrorReason(w http.ResponseWriter, statusCode int, errorType app.ErrorType, reason string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(newErrorResponse(errorType, reason))
	if err != nil {
		log.Printf("Failed to encode the response: %s\n", err)
	}
}

// RespondJson responds with OK status code and JSON as body and content type.
func RespondJson(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(body)
}
