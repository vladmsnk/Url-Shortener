package main

import (
	"log"
	"vladmsnk/urlshort/config"
	"vladmsnk/urlshort/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	app.Run(cfg)
}
