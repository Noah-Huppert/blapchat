package db

import "github.com/jinzhu/gorm"

// Picture is the resource stores information related to images sent by users
// on the platform
type Picture struct {
	gorm.Model

	// User will be populated with the User who owns the Picture
	User User

	// UserID is the id of the user who owns the picture
	UserID uint `gorm:"not null"`

	// B2BucketID is the id of the bz2 bucket which holds picture data
	B2BucketID string `gorm:"not null"`

	// B2FileID is the id of the bz2 file which holds picture data
	B2FileID string `gorm:"not null"`

	// B2FileName is the name of the bz2 file which holds picture data
	B2FileName string `gorm:"not null"`

	// B2FileURL is the url at which the raw bz2 file data can be accessed
	B2FileURL string `gorm:"not null"`

	// Text is the string message present in the picture. Set to nil if
	// none is present.
	Text string
}
