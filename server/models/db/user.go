package db

import (
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model

    Username string

    FirstName string
    LastName string

    Email string
    EmailVerified bool

    BzProfilePictureId string
}

// NewUser creates and returns a new user with the specified parameters
func NewUser(Username string,
            FirstName string,
            LastName string,
            Email string,
            EmailVerified bool,
            BzProfilePictureId string) User {
    return User{
        Username: Username,
        FirstName: FirstName,
        LastName: LastName,
        Email: Email,
        EmailVerified: EmailVerified,
        BzProfilePictureId: BzProfilePictureId,
    }
}
