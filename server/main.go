package main

import (
    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"

    "./models"
    tables "./models/db"
)

func main() {
    // Setup
    // -- -- Load config
    config := models.NewConfig("localhost", "user", "password", "snap-but-doesnt-suck")

    // -- -- Connect to db
    db, err := gorm.Open("postgres", config.GetDbConnOpts("sslmode=disable"))
    if err != nil {
        fmt.Errorf("Failed to connect to database: %s", err)
    }
    defer db.Close()

    // -- -- Setup DB tables
    tables.Load(db)
}
