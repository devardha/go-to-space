package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB //database

// GetEnv - get variables from .env
func GetEnv(name string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot load .env file")
	}

	return os.Getenv(name)
}

// ConnectToDb - initializes the ORM and Connection to the postgres DB
func ConnectToDb() {
	dsn := GetEnv("CONNECTION_STRING")
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	if conn != nil {
		fmt.Println("Database connected")
	}

	db = conn
	db.Debug().AutoMigrate() //Database migration
}

// GetDB - get database object
func GetDB() *gorm.DB {
	return db
}

func main() {
	ConnectToDb()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	port := os.Getenv("PORT")
	//port := "5000"

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Serving template & static files
	r.LoadHTMLGlob("./web/templates/*")
	r.Static("/static", "./web/static")

	// Frontend Routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "maintemplate.gohtml", gin.H{"status": "OK"})
	})

	// Api Routes

	// Run server
	fmt.Println("Server running")
	r.Run(":" + port)
}
