package main

import (
	"fmt"
	"log"
	"os"

	"github.com/devardha/gotospace/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectToDb()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	port := os.Getenv("PORT")
	fmt.Println("PORT: " + port)

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// Serving template & static files
	r.LoadHTMLGlob("./web/templates/*")
	r.Static("/static", "./web/static")

	// Frontend Routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "maintemplate.gohtml", gin.H{"status": "200"})
	})

	// Run server
	fmt.Println("Server running")
	r.Run(":" + port)
}
