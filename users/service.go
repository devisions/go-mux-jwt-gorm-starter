package users

// UserService
type UserService interface {
	Authenticate(email, password string) (*User, error)
	Uninit()
}

func NewUserService(dbConnInfo string) (UserService, error) {
	userStore, err := newUserStoreGorm(dbConnInfo)
	if err != nil {
		return nil, err
	}
	return &userService{
		userStore,
	}, nil
}

// Implementation of the UserService.
type userService struct {
	userStore UserStore
}

func (us *userService) Authenticate(email, password string) (*User, error) {
	return nil, nil
}

func (us *userService) Uninit() {
	// TODO
}
