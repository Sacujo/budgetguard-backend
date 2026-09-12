package config

import (
	"errors"
	"os"
)

type Config struct {
	Port  string
	DBUrl string
}

func Load() (*Config, error) {
	url := os.Getenv("DB_URL")
	if url == "" {
		return nil, errors.New("DB_URL is required, example: postgres://user:pass@localhost:5432/budgetguard")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		Port:  port,
		DBUrl: url,
	}, nil
}
