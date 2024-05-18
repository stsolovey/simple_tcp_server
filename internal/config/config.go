package config

import (
	"errors"
	"os"
)

type Config struct {
	Host string
	Port string
}

var (
	errMissingPort = errors.New("missing port variable")
	errMissingHost = errors.New("missing host variable")
)

func New() (*Config, error) {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	switch {
	case port == "":
		return nil, errMissingPort
	case host == "":
		return nil, errMissingHost
	}

	return &Config{
			Host: host,
			Port: port,
		},
		nil
}
