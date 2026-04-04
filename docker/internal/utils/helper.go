package utils

import "os"

func GetEnv(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return os.Getenv(value)
}
