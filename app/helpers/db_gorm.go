package helpers

import (
	"log"

	"github.com/devisions/go-mux-jwt-gorm-starter/app"
	"gorm.io/gorm"
)

// FirstRecord queries using the provided gorm.DB to get the first record returned
// and place it into dst. If nothing found, app.ErrNotFound will be returned.
func FirstRecord(db *gorm.DB, dst interface{}) error {

	err := db.First(dst).Error
	if err != nil {
		log.Printf("Error getting first record: %s\n", err)
		if err == gorm.ErrRecordNotFound {
			return app.ErrNotFound.AsError()
		}
	}
	return err
}
