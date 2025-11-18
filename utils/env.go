package utils

import (
	"os"
	"strconv"
	"strings"
)

// GetEnv gets an environment variable with a fallback value
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

// GetEnvAsInt gets an environment variable as integer with fallback
func GetEnvAsInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return fallback
}

// GetEnvAsBool gets an environment variable as boolean with fallback
func GetEnvAsBool(key string, fallback bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return fallback
}

// GetAllowedOrigins gets CORS allowed origins from environment
func GetAllowedOrigins() []string {
	origins := GetEnv("ALLOWED_ORIGINS", "*")
	if origins == "*" {
		return []string{"*"}
	}
	return strings.Split(origins, ",")
}
