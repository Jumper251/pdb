package utils

import "os"

type Config struct {
	AuthUser     string
	AuthPassword string
}

func GetConfig() *Config {
	return &Config{
		AuthUser:     GetenvOrDefault("AUTH_USER", "test"),
		AuthPassword: GetenvOrDefault("AUTH_PASSWORD", "1234"),
	}
}

func GetenvOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
