package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB = ConnectDB()

func ConnectDB() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("failed to load `.env` file error: %s", err.Error())
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=9191 TimeZone=Asia/Dushanbe",
		dbHost, dbUser, dbPass, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("failed to connet database: %s", err.Error())
	}

	return db
}

func Disconnect(db *gorm.DB) {
	_db, err := db.DB()
	if err != nil {
		logrus.Fatalf("failed to disconnecting with database: %s", err.Error())
	}

	_db.Close()
}
