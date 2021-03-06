package users

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID    int    `gorm:"primary_key" json:"id"`
	Email string `gorm:"not null;unique_index" json:"email"`
	Name  string `json:"name"`
	Hash  string `json:"-"` // skipped on json marshalling
}

func (u User) checkPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Hash), []byte(password))
	return err == nil
}
