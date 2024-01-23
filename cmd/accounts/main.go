package main

import (
	"log"

	"github.com/indigowar/delivery/internal/accounts/app"
	"github.com/indigowar/delivery/internal/accounts/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Accounts service has invalid configuration: %s", err.Error())
	}
	app.Run(cfg)
}
