package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default() // Initialize Gin router with default gin config

	// define ping route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Run the server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
