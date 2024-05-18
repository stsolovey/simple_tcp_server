package config

import (
	"errors"
	"os"
)

type Config struct {
	Port string
}

var errMissingPort = errors.New("missing port variable")

func New() (*Config, error) {
	port := os.Getenv("PORT")

	if port == "" {
		return nil, errMissingPort
	}

	return &Config{
			Port: port,
		},
		nil
}
