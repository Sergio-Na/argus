package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Warning: Error loading .env file: %v\n", err)
	}

	serverAddress, err := getEnv("SERVER_ADDRESS")
	if err != nil {
		return nil, fmt.Errorf("error loading SERVER_ADDRESS: %w", err)
	}

	return &Config{
		ServerAddress: serverAddress,
	}, nil
}

func getEnv(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists {
		return "", fmt.Errorf("environment variable %s not found", key)
	}
	if value == "" {
		return "", fmt.Errorf("environment variable %s is empty", key)
	}
	return value, nil
}
