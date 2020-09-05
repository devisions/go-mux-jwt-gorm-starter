package users

// UserService
type UserService interface {
	Authenticate(email, password string) (*User, error)
}

func NewUserService(store UserStore) (UserService, error) {
	return &userService{
		userStore: store,
	}, nil
}

// Implementation of the UserService.
type userService struct {
	userStore UserStore
}

func (us *userService) Authenticate(email, password string) (*User, error) {
	return nil, nil
}
