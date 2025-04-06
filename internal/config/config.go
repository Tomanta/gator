package config

import (
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	db_url string
	current_user_name string
}

func Read() (Config, error) {
	return Config{}, nil
}

func getConfigFilePath() (string, error) {
	return os.UserHomeDir()
}

func write(cfg Config) error {
	return nil
}

func (cfg *Config) SetUser(user_name string) error {
	cfg.current_user_name = user_name
	return write(*cfg) // If there is an error, just ripple it up for now.
}