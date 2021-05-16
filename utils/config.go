package utils

import "os"

type Config struct {
	AuthUser     string
	AuthPassword string
	Mode         string
}

func (c Config) IsRelease() bool {
	return c.Mode == "release"
}

func GetConfig() *Config {
	return &Config{
		AuthUser:     GetenvOrDefault("AUTH_USER", "test"),
		AuthPassword: GetenvOrDefault("AUTH_PASSWORD", "1234"),
		Mode:         GetenvOrDefault("GIN_MODE", "development"),
	}
}

func GetenvOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
