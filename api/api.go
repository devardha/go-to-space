package api

import (
	"github.com/devardha/gotospace/api/handlers"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes (grouping api)
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("api")
	{
		// Galaxy routes
		api.GET("/galaxies", handlers.GetGalaxies())
		api.GET("/galaxies/:id", handlers.GetGalaxy())

		api.POST("/galaxies", handlers.AddGalaxy())
		api.DELETE("/galaxies/:id", handlers.DeleteGalaxy())
		api.PUT("/galaxies/:id", handlers.UpdateGalaxy())

		// Star routes
		api.GET("/stars", handlers.GetStars())
		api.GET("/stars/:id", handlers.GetStar())

		api.POST("/stars", handlers.AddStar())
		api.DELETE("/stars/:id", handlers.DeleteStar())
		api.PUT("/stars/:id", handlers.UpdateStar())
	}
}
