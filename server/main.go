package main

import (
    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"

    "github.com/Noah-Huppert/blapchat/server/models"
    tables "github.com/Noah-Huppert/blapchat/server/models/db"
    "net/http"
    "github.com/Noah-Huppert/blapchat/server/transports"
    "os"
    "github.com/Noah-Huppert/blapchat/server/handlers"
)

func main() {
    // Config
    err, config := models.LoadConfigFile(models.DEFAULT_CONFIG_PATH)
    if err != nil {
        fmt.Printf("Error loading settings.config.json, %s", err.Error())
        return
    }

    // -- -- DB
    db, err := gorm.Open("postgres", config.GetDbConnOpts("sslmode=disable"))
    if err != nil {
        fmt.Errorf("Failed to connect to database: %s", err)
    }
    defer db.Close()

    tables.Load(db)

    // -- -- HTTP
    httpMux := http.NewServeMux()
    httpTransport := transports.NewHttpTransport(httpMux)
    httpTransport.RegisterHandler(handlers.NewTestHandler())

    httpPort := fmt.Sprintf(":%d", config.Port)

    fmt.Printf("Http transport starting on \"%s\"\n", httpPort)
    if err := http.ListenAndServe(httpPort, httpMux); err != nil {
        fmt.Printf("Error starting Http transport on \"%s\", error: \"%s\"\n", httpPort, err.Error())
        os.Exit(1)
    }
}
