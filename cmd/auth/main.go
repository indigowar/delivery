package main

import (
	"log"

	"github.com/indigowar/delivery/internal/auth/app"
	"github.com/indigowar/delivery/internal/auth/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Auth service has invalid configuration: %s", err.Error())
	}

	app.Run(cfg)
}
