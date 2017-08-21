package db

import "github.com/jinzhu/gorm"

type Picture struct {
    gorm.Model

    User User
    UserId uint `gorm:"not null"`

    B2BucketId string `gorm:"not null"`
    B2FileId string `gorm:"not null"`
    B2FileName string `gorm:"not null"`
    B2FileUrl string `gorm:"not null"`

    Text string
}
