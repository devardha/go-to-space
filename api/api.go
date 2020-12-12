package api

import (
	"github.com/devardha/gotospace/api/handlers"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes (grouping api)
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("api")
	{
		api.GET("/galaxies", handlers.GetGalaxies())
		api.GET("/galaxies/:id", handlers.GetGalaxy())
		api.GET("/stars", handlers.GetStars())
		api.GET("/stars/:id", handlers.GetStar())
	}
}
