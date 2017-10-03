package db

import (
	"github.com/jinzhu/gorm"
)

// User is the resource which holds information about individual accounts on
// the BlapChat platform
type User struct {
	gorm.Model

	// FirstName is the user's first name, sometimes called the: given name
	FirstName string `gorm:"not null"`

	// LastName is the user's last name, sometimes called the: family name
	LastName string `gorm:"not null"`

	// Username is the screen name which the user goes by
	Username string

	// ProfilePictureID is the id of the Picture which holds the user's
	// profile picture
	ProfilePictureID uint `gorm:"not null"`

	// PasswordHash holds the hash of the user's password
	PasswordHash string `json:"-" gorm:"not null"`
}
