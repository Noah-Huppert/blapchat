package db

import (
    "github.com/jinzhu/gorm"
    "time"
)

type DirectMessage struct {
    gorm.Model

    Sender User `gorm:"ForeignKey:SenderUserId"`
    SenderUserId uint `gorm:"not null"`

    Recipient User `gorm:"ForeignKey:RecipientUserId"`
    RecipientUserId uint `gorm:"not null"`

    Picture Picture
    PictureId uint `gorm:"not null"`

    Text string `gorm:"not null"`
    Length int `gorm:"not null"`

    SentAt time.Time `gorm:"not null"`
    ViewedAt time.Time
}
