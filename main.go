package main

import (
	"github.com/gbmoraes-dev/gopportunities/config"
	"github.com/gbmoraes-dev/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	// Initialize logger
	logger = config.GetLogger("main")

	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
		return
	}

	// Initialize the routes
	router.Initialize()
}