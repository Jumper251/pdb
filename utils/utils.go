package utils

import "os"

func EnsureDirectoryExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.Mkdir(path, os.ModeDir|0755)
	}
}
