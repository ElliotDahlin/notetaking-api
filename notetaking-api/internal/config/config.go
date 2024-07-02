package config

import (
    "os"
)

type Config struct {
    ServerAddress string
}

func LoadConfig() (*Config, error) {
    return &Config{
        ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
    }, nil
}

func getEnv(key, fallback string) string {
    value := os.Getenv(key)
    if value == "" {
        return fallback
    }
    return value
}
