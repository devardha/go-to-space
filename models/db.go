package models

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB //database

// GetEnv - get variables from .env
func GetEnv(name string) string {
	godotenv.Load(".env")

	return os.Getenv(name)
}

// ConnectToDb - initializes the ORM and Connection to the postgres DB
func ConnectToDb() {
	dsn := GetEnv("CONNECTION_STRING")
	conn, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if conn != nil {
		fmt.Println("Database connected")
	}

	db = conn
	db.Debug().AutoMigrate(&Star{}, &Galaxy{}) //Database migration
}

// GetDB - get database object
func GetDB() *gorm.DB {
	return db
}
