package users

// UserService
type UserService interface {

	// embedded
	UserStore

	Authenticate(email, password string) (*User, error)
}

func NewUserService(dbConnInfo string) (UserService, error) {
	us, err := newUserStoreGorm(dbConnInfo)
	if err != nil {
		return nil, err
	}
	return &userService{
		UserStore: us,
	}, nil
}

// Implementation of the UserService.
type userService struct {
	UserStore
}

func (us *userService) Authenticate(email, password string) (*User, error) {
	return nil, nil
}
