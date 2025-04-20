package main

import (
	"github.com/olajoao/solidari-api/config"
	"github.com/olajoao/solidari-api/router"
)

var logger *config.Logger

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("Main Initialize error: %v", err)
		return
	}

	router.Initialize()
}
