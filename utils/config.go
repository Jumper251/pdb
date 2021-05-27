package utils

import "os"

type Config struct {
	AuthUser     string
	AuthPassword string
	Mode         string
	FtpUrl       string
	FtpUser      string
	FtpPassword  string
}

func (c Config) IsRelease() bool {
	return c.Mode == "release"
}

func GetConfig() *Config {
	return &Config{
		AuthUser:     GetenvOrDefault("AUTH_USER", "test"),
		AuthPassword: GetenvOrDefault("AUTH_PASSWORD", "1234"),
		Mode:         GetenvOrDefault("GIN_MODE", "development"),
		FtpUrl:       GetenvOrDefault("FTP_URL", ""),
		FtpUser:      GetenvOrDefault("FTP_USER", ""),
		FtpPassword:  GetenvOrDefault("FTP_PASSWORD", ""),
	}
}

func GetenvOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
