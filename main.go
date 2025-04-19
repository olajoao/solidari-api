package main

import (
	"fmt"

	"github.com/olajoao/solidari-api/config"
	"github.com/olajoao/solidari-api/router"
)

func main() {
	// Initialize Configs
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}

	router.Initialize()
}
