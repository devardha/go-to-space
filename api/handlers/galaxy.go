package handlers

import (
	"fmt"

	"github.com/devardha/gotospace/models"

	"github.com/gin-gonic/gin"
)

// GetGalaxies (get all galxy)
func GetGalaxies() gin.HandlerFunc {
	return func(c *gin.Context) {
		database := models.GetDB()
		var galaxies models.Galaxies
		galaxyResult := database.Find(&galaxies)

		if galaxyResult == nil {
			fmt.Println("Query failed")
		}

		c.JSON(200, galaxies)
	}
}

// GetGalaxy (get galaxy by id)
func GetGalaxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var galaxies models.Galaxies
		database := models.GetDB()
		galaxyResult := database.First(&galaxies, id)

		if galaxyResult == nil {
			fmt.Println("")
		}

		c.JSON(200, galaxies)
	}
}
