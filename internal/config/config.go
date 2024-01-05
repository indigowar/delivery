package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Postgres struct {
		Host string
		Port int
		Db   string

		User     string
		Password string
	}
}

// LoadConfig - loads the config from ENV or other configuration files
func LoadConfig() (*Config, error) {
	var cfg Config

	cfg.Postgres.Host = os.Getenv("POSTGRES_HOST")

	port := os.Getenv("POSTGRES_PORT")
	parsed, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse the configuration(postgres port) - %w", err)
	}
	cfg.Postgres.Port = parsed

	cfg.Postgres.Db = os.Getenv("POSTGRES_DB")

	cfg.Postgres.User = os.Getenv("POSTGRES_USER")
	cfg.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")

	return &cfg, nil
}
