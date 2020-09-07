package users

import (
	"log"
	"os"
	"time"

	"github.com/devisions/go-mux-jwt-gorm-starter/api/rest/config"
	"github.com/devisions/go-mux-jwt-gorm-starter/api/rest/tokens"
	"github.com/devisions/go-mux-jwt-gorm-starter/app"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// UserService
type UserService interface {
	Authenticate(email, password string) (*tokens.JWTToken, error)
	Register(name, email, password string) (*User, error)
	GetByID(id uint) (*User, error)
}

func NewUserService(store UserStore) UserService {
	return &userService{
		userStore: store,
	}
}

// Implementation of the UserService.
type userService struct {
	userStore UserStore
}

func (us *userService) Authenticate(email, password string) (*tokens.JWTToken, error) {
	user, err := us.userStore.GetByEmail(email)
	if err != nil {
		log.Printf("Error at authenticate: %s\n", err)
		return nil, err
	}
	if user.checkPassword(password) {
		return generateJWT(*user)
	}
	return nil, app.ErrInvalidCreds.AsError()
}

func (us *userService) Register(name, email, password string) (*User, error) {
	user := User{
		Name:  name,
		Email: email,
		Hash:  us.hashPassword(password),
	}
	if err := us.userStore.Create(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *userService) GetByID(id uint) (*User, error) {
	return us.userStore.GetByID(id)
}

func (us *userService) hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes)
}

// generateJWT generates a signed JWT with user details as claims.
func generateJWT(user User) (*tokens.JWTToken, error) {
	signingKey := []byte(os.Getenv(config.JwtSignKey))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(1 * time.Hour).Unix(),
		"user_id": int(user.ID),
		"name":    user.Name,
		"email":   user.Email,
	})
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		log.Printf("Error while generating jwt: %s\n", err)
	}
	return &tokens.JWTToken{
		Token: tokenString,
	}, err
}
