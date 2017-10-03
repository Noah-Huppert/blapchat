package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

// DirectMessage is a resource which stores information about a message sent
// from one user to another.
type DirectMessage struct {
	gorm.Model

	// Sender will be populated with the User who sent and owns the message
	Sender User `gorm:"ForeignKey:SenderUserID"`

	// SenderUserID is the id of the user who owns the message
	SenderUserID uint `gorm:"not null"`

	// Recipient will be populated with the User who the message is sent to
	Recipient User `gorm:"ForeignKey:RecipientUserID"`

	// RecipientUserID is the id of the user who the message is sent to
	RecipientUserID uint `gorm:"not null"`

	// Picture will be populated with the Picture included with the message
	// , if no picture is included it will be nil
	Picture Picture

	// PictureID is the id of the included picture, if no picture is
	// included it will be -1
	PictureID uint `gorm:"not null"`

	// Text is the text content of the message being sent to the other user
	// , if no text is included it is set to ""
	Text string `gorm:"not null"`

	// Length is the time in seconds the recipient user can view the
	// message. Set to -1 if infinite view length.
	Length int `gorm:"not null"`

	// SentAt is the time the message was sent
	SentAt time.Time `gorm:"not null"`

	// ViewedAt is the time the recipient user view the message. Set to nil
	// if not viewed yet
	ViewedAt time.Time
}
