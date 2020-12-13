package handlers

import (
	"fmt"

	"github.com/devardha/gotospace/models"

	"github.com/gin-gonic/gin"
)

// GalaxyInput (galaxy input model)
type GalaxyInput struct {
	Name          string `json:"name" gorm:"autoIncrement:true;primaryKey"`
	Constellation string `json:"constellation" binding:"required"`
	Type          string `json:"type" binding:"required"`
}

// GetGalaxies (get all galxy) - GET
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

// GetGalaxy (get galaxy by id) - GET
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

// AddGalaxy (add new galaxy) - POST
func AddGalaxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		database := models.GetDB()

		var input GalaxyInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"data": err.Error()})
			return
		}

		galaxy := models.Galaxy{Name: input.Name, Constellation: input.Constellation, Type: input.Type}
		database.Create(&galaxy)

		c.JSON(200, gin.H{"data": galaxy})
	}
}

// UpdateGalaxy (update galaxy by id) - PUT
func UpdateGalaxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		database := models.GetDB()

		var input GalaxyInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.AbortWithStatus(400)
			return
		}

		var galaxy models.Galaxy
		if err := database.Where("id = ?", id).First(&galaxy).Error; err != nil {
			c.AbortWithStatus(404)
			return
		}

		database.Model(&galaxy).Updates(models.Galaxy{
			Name:          input.Name,
			Constellation: input.Constellation,
			Type:          input.Type,
		})
		c.JSON(200, gin.H{"msg": "Data updated successfully"})
	}
}

// DeleteGalaxy (delete galaxy by id) - DELETE
func DeleteGalaxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		database := models.GetDB()
		id := c.Param("id")
		var galaxy models.Galaxy

		result := database.First(&galaxy, id)

		if result == nil {
			c.JSON(400, gin.H{"data": "Galaxy not found"})
		}

		database.Delete(&galaxy)
		c.JSON(200, gin.H{"data": "Galaxy deleted succesfully"})
	}
}
