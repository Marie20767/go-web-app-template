package config

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbTimeout int
	DbURL     string
	LogLevel  slog.Level
	Port      string
}

var logLevelMap = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

func ParseEnv() (*Config, error) {
	// Ignore error because in production there will be no .env file, env vars will be passed
	// in at runtime via docker run command/docker compose
	_ = godotenv.Load()

	envVars := map[string]string{
		"DATABASE_URL": "",
		"LOG_LEVEL":    "",
		"SERVER_PORT":  "",
	}

	for key := range envVars {
		value := os.Getenv(key)
		if value == "" {
			return nil, fmt.Errorf("%s environment variable is not set", key)
		}
		envVars[key] = value
	}

	logLevel, ok := logLevelMap[envVars["LOG_LEVEL"]]
	if !ok {
		return nil, errors.New("LOG_LEVEL should be one of debug|info|warning|error")
	}

	return &Config{
		DbURL:    envVars["DATABASE_URL"],
		LogLevel: logLevel,
		Port:     envVars["SERVER_PORT"],
	}, nil
}
