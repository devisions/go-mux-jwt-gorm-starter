package rest

import (
	"fmt"
	"github.com/devisions/go-mux-jwt-gorm-starter/api/rest/responses"
	"github.com/devisions/go-mux-jwt-gorm-starter/app"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			responses.RespondWithError(w, http.StatusUnauthorized, app.UnauthorizedError)
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := VerifyToken(tokenString)
		if err != nil {
			responses.RespondWithErrorReason(w, http.StatusUnauthorized, app.JWTValidationError, err.Error())
			return
		}
		userId := fmt.Sprintf("%v", claims.(jwt.MapClaims)["user_id"])
		r.Header.Set("userId", userId)
		next.ServeHTTP(w, r)
	})
}
