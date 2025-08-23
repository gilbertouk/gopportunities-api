package router

import (
	"net/http"

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
		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GET openings",
			})
		})

		v1.POST("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusCreated, gin.H{
				"message": "POST openings",
			})
		})

		v1.DELETE("/openings/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "DELETE opening " + ctx.Param("id"),
			})
		})

		v1.PUT("/openings/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "PUT opening " + ctx.Param("id"),
			})
		})

		v1.GET("/openings/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "GET opening " + ctx.Param("id"),
			})
		})
	}

}
