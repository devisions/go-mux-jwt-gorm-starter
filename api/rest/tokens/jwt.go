package tokens

import (
	"os"

	"github.com/devisions/go-mux-jwt-gorm-starter/api/rest/config"
	"github.com/dgrijalva/jwt-go"
)

type JWTToken struct {
	Token string `json:"token"`
}

// VerifyJWTToken parses the JWT token to validate if it is correctly signed
// and returns the claims that are included.
func VerifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte(os.Getenv(config.JwtSignKey))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, nil
}
