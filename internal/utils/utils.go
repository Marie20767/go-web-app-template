package utils

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  string
	DbURL string
}

func ParseEnv() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	dbURL := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")

	if dbURL == "" || port == "" {
		return nil, errors.New("not all environment variables are set")
	}

	return &Config{
		Port:  port,
		DbURL: dbURL,
	}, nil

}
