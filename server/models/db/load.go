package db

import (
    "github.com/jinzhu/gorm"
)

// Load calls the GORM DB.AutoMigrate method on all database structs in the db package
// db: GORM database to migrate
func Load(db *gorm.DB) {
    db.AutoMigrate(&User{})
    db.AutoMigrate(&Picture{})
    db.AutoMigrate(&DirectMessage{})
}
