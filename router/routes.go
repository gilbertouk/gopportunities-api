package router

import (
	"net/http"

	"github.com/gilbertouk/gopportunities-api/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	// Ping route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/api/v1")
	{
		// Define your v1 routes here

		// opening routes
		v1.GET("/openings", handler.ListOpeningsHandler)
		v1.POST("/openings", handler.CreateOpeningHandler)
		v1.DELETE("/openings/:id", handler.DeleteOpeningHandler)
		v1.PUT("/openings/:id", handler.UpdateOpeningHandler)
		v1.GET("/openings/:id", handler.ShowOpeningHandler)
	}

}
