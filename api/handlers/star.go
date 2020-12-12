package handlers

import (
	"fmt"

	"github.com/devardha/gotospace/models"

	"github.com/gin-gonic/gin"
)

// GetStars (get all stars)
func GetStars() gin.HandlerFunc {
	return func(c *gin.Context) {
		database := models.GetDB()
		var stars models.Stars
		starResult := database.Find(&stars)
		if starResult == nil {
			fmt.Println("Query failed")
		}

		c.JSON(200, stars)
	}
}

// GetStar (get star by id)
func GetStar() gin.HandlerFunc {
	return func(c *gin.Context) {
		database := models.GetDB()
		var stars models.Star
		id := c.Param("id")
		starResult := database.First(&stars, id)

		if starResult == nil {
			fmt.Print("Querry failed")
		}

		c.JSON(200, stars)

	}
}
