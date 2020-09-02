package rest

import (
	"github.com/devisions/go-mux-jwt-gorm-starter/api/rest/config"
	"github.com/devisions/go-mux-jwt-gorm-starter/user/domain"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
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

// generateJWT generates a signed JWT with user details as claims.
func generateJWT(user domain.User) (JWTToken, error) {
	signingKey := []byte(os.Getenv(config.JwtSignKey))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(1 * time.Hour).Unix(),
		"user_id": int(user.ID),
		"name":    user.Name,
		"email":   user.Email,
	})
	tokenString, err := token.SignedString(signingKey)
	return JWTToken{tokenString}, err
}