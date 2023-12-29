package main

import (
	"log"

	"github.com/indigowar/delivery/internal/app"
	"github.com/indigowar/delivery/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Project configuration is invalid: %w", err)
	}
	app.Run(cfg)
}
