package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// to allow mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

// DB global handle to database
var DB *gorm.DB

// LoadEnv to load or init variables from .env
func loadEnv() {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load(".env")
	}

	if err != nil {
		panic("Error loading .env file")
	}
}

// ConnectDB - setup DB connection
func ConnectDB() {
	loadEnv()
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", username, password, dbHost, dbPort, dbName)
	db, err := gorm.Open("mysql", connectionStr)

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Car{})

	DB = db
}
