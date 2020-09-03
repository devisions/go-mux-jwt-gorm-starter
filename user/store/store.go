package store

import (
	"github.com/devisions/go-mux-jwt-gorm-starter/user/domain"
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
	GetByID(int uint) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)

	Create(user *domain.User) error
	Update(user *domain.User) error
	Delete(id uint) error

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
func newUserStoreGorm(connectionInfo string) (*userStoreGorm, error) {

	// TODO: have dsn params in external config
	dsn := "user=postgres password=postgres DB.name=starter port=54325 sslmode=disable"
	cfg := postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}
	db, err := gorm.Open(postgres.New(cfg), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.Logger = db.Logger.LogMode(logger.Info)
	return &userStoreGorm{db: db}, nil
}
