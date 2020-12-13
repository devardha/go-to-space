package handlers

import (
	"github.com/devardha/gotospace/models"

	"github.com/gin-gonic/gin"
)

// StarInput (input model)
type StarInput struct {
	Name          string  `json:"name" binding:"required"`
	Constellation string  `json:"constellation" binding:"required"`
	Luminosity    float32 `json:"luminosity" binding:"required"`
	Temperature   float32 `json:"temperature" binding:"required"`
	Distance      float32 `json:"distance_ly" binding:"required"`
}

// GetStars (get all stars) - GET
func GetStars() gin.HandlerFunc {
	return func(c *gin.Context) {
		database := models.GetDB()
		var stars models.Stars
		starResult := database.Find(&stars)
		if starResult == nil {
			c.JSON(400, gin.H{"data": "Stars not found"})
		}

		c.JSON(200, stars)
	}
}

// GetStar (get star by id) - GET
func GetStar() gin.HandlerFunc {
	return func(c *gin.Context) {
		database := models.GetDB()
		var stars models.Star
		id := c.Param("id")
		starResult := database.First(&stars, id)

		if starResult == nil {
			c.JSON(400, gin.H{"data": "Stars not found"})
		}

		c.JSON(200, stars)

	}
}

// AddStar (add new star) - POST
func AddStar() gin.HandlerFunc {
	return func(c *gin.Context) {
		database := models.GetDB()

		var input StarInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"data": err.Error()})
			return
		}

		star := models.Star{Name: input.Name, Constellation: input.Constellation, Luminosity: input.Luminosity, Temperature: input.Temperature, Distance: input.Distance}
		database.Create(&star)

		c.JSON(200, gin.H{"data": star})
	}
}

// UpdateStar (update star data by ud) - PUT
func UpdateStar() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		database := models.GetDB()

		var input StarInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.AbortWithStatus(400)
			return
		}

		var star models.Star
		if err := database.Where("id = ?", id).First(&star).Error; err != nil {
			c.AbortWithStatus(404)
			return
		}

		database.Model(&star).Updates(models.Star{
			Name:          input.Name,
			Constellation: input.Constellation,
			Luminosity:    input.Luminosity,
			Temperature:   input.Temperature,
			Distance:      input.Distance,
		})
		c.JSON(200, gin.H{"msg": "Data updated successfully"})
	}
}

// DeleteStar (delete star by id) - DELETE
func DeleteStar() gin.HandlerFunc {
	return func(c *gin.Context) {
		database := models.GetDB()
		id := c.Param("id")
		var star models.Star

		result := database.First(&star, id)

		if result == nil {
			c.JSON(400, gin.H{"data": "Star not found"})
		}

		database.Delete(&star)
		c.JSON(200, gin.H{"msg": "Star deleted successfully"})

	}
}
