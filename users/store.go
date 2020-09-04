package users

import (
	"log"

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

	// Create(user *User) error
	// Update(user *User) error
	// Delete(id uint) error

	// Close is used for closing the connection(s) to the store (database).
	Close()

	// AutoMigrate is a helper method used for database migration
	AutoMigrate() error
}

// An implementation of `UserStore` interface, used internally.
type userStoreGorm struct {
	db *gorm.DB
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
	log.Println("Connected to database.")
	return &userStoreGorm{db}, nil
}

func (us *userStoreGorm) GetByID(int uint) (*User, error) {
	return nil, nil
}

func (us *userStoreGorm) GetByEmail(email string) (*User, error) {
	return nil, nil
}

func (us *userStoreGorm) Close() {

}

func (us *userStoreGorm) AutoMigrate() error {
	return nil
}
