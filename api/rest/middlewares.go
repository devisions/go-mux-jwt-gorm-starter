package rest

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/devisions/go-mux-jwt-gorm-starter/api/rest/responses"
	"github.com/devisions/go-mux-jwt-gorm-starter/api/rest/tokens"
	"github.com/devisions/go-mux-jwt-gorm-starter/app"
	"github.com/dgrijalva/jwt-go"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			responses.RespondJsonWithUnauthorizedError(w)
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := tokens.VerifyToken(tokenString)
		if err != nil {
			responses.RespondJsonWithErrorReason(w, http.StatusUnauthorized, app.ErrJWTValidation, err.Error())
			return
		}
		userId := fmt.Sprintf("%v", claims.(jwt.MapClaims)["user_id"])
		r.Header.Set("userId", userId)
		next.ServeHTTP(w, r)
	})
}
