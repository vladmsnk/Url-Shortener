package main

import (
	"log"

	"vladmsnk/taskrec/config"
	"vladmsnk/taskrec/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
