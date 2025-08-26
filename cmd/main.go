package main

import (
	"github.com/gilbertouk/gopportunities-api/config"
	"github.com/gilbertouk/gopportunities-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	logger.Info("Go Opportunities Rest API")

	// Initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	// Initialize the router
	router.Initialize()
}
