package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/labstack/echo"

	"github.com/Noah-Huppert/blapchat/server/models"
	tables "github.com/Noah-Huppert/blapchat/server/models/db"
	//"github.com/NaySoftware/go-fcm"
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

	// -- -- FCM
	//fcmClient := fcm.NewFcmClient(config.FCMServerKey)

	// -- -- HTTP
	e := echo.New()

	err = e.Start(fmt.Sprintf(":%d", config.Port))
	if err != nil {
		fmt.Errorf("Error starting http server on port %d, %s", config.Port, err.Error())
	}
}
