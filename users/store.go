package users

import (
	"log"

	"github.com/devisions/go-mux-jwt-gorm-starter/app/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// UserStore is used for interacting with the persistence layer for users,
// and it is the contract (interface) to used by outside package components.
// For all `Get_` methods, it either returns the user that is found and a nil error,
// or an error that is either defined by the `models` package (such as `ErrNotFound`)
// or another, more low level error.
type UserStore interface {
	GetByID(int uint) (*User, error)
	GetByEmail(email string) (*User, error)

	Create(user *User) error
	// Update(user *User) error
	// Delete(id uint) error

	// Close is used for closing the connection(s) to the store (database).
	Close()

	// AutoMigrate is a helper method used for database migration
	Migrate() error
}

// An implementation of `UserStore` interface, used internally.
type userStoreGorm struct {
	db *gorm.DB
}

func NewUserStore(dbConnInfo string) (UserStore, error) {
	return newUserStoreGorm(dbConnInfo)
}

// Internal constructor of a userStoreGorm instance.
func newUserStoreGorm(dbConnInfo string) (*userStoreGorm, error) {

	cfg := postgres.Config{
		DSN:                  dbConnInfo,
		PreferSimpleProtocol: true,
	}
	db, err := gorm.Open(postgres.New(cfg), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.Logger = db.Logger.LogMode(logger.Info)
	log.Println("Connected to user store database.")
	return &userStoreGorm{db}, nil
}

func (us *userStoreGorm) GetByID(id uint) (*User, error) {
	var user User
	db := us.db.Where("id = ?", id)
	if err := helpers.FirstRecord(db, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *userStoreGorm) GetByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
	if err := helpers.FirstRecord(db, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *userStoreGorm) Close() {
	// TODO: is there a .Close() in GORM v2?
}

func (us *userStoreGorm) Create(user *User) error {
	return us.db.Create(user).Error
}

func (us *userStoreGorm) Migrate() error {
	log.Println("Running database migration now.")
	return us.db.AutoMigrate(&User{})
}
