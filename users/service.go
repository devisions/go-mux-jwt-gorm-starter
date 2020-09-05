package users

import (
	"golang.org/x/crypto/bcrypt"
)

// UserService
type UserService interface {
	Authenticate(email, password string) (*User, error)
	Register(name, email, password string) (*User, error)
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

func (us *userService) Authenticate(email, password string) (*User, error) {
	return nil, nil
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

func (us *userService) hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes)
}
