package config

type Config struct{}

// LoadConfig - loads the config from ENV or other configuration files
func LoadConfig() (*Config, error) {
	return &Config{}, nil
}
