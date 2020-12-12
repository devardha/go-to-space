package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	port := os.Getenv("PORT")

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
