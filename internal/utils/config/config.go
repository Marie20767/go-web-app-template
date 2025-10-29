package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	DbURL     string
	DbTimeout int
}

func ParseEnv() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	envVars := map[string]*string{
		"PORT":       nil,
		"DB_URL":     nil,
		"DB_TIMEOUT": nil,
	}

	for key := range envVars {
		value := os.Getenv(key)
		if value == "" {
			return nil, errors.New("not all environment variables are set")
		}
		envVars[key] = &value
	}

	dbTimeout, err := strconv.Atoi(*envVars["DB_TIMEOUT"])
	if err != nil {
		return nil, err
	}

	return &Config{
		Port:      *envVars["PORT"],
		DbURL:     *envVars["DB_URL"],
		DbTimeout: dbTimeout,
	}, nil
}
