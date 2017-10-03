package db

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`

	Username         string
	ProfilePictureId uint `gorm:"not null"`

	PasswordHash string `json:"-" gorm:"not null"`
}
