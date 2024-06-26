package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hryhorskyi/gin-app/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	var err error
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName, dbPassword)

	DB, err = gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB.LogMode(true)

	DB.AutoMigrate(&models.Subscription{})

	DB.Callback().Update().Register("gorm:update_time_stamp", updateTimestampForUpdateCallback)
}

func updateTimestampForUpdateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		scope.SetColumn("UpdatedAt", time.Now())
	}
}
