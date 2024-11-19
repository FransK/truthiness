package env

import (
	"os"
	"strconv"
)

// GetString(key, fallback string) string
// Looks for a key in Environment Variables
// If found, returns value as a string
// If not found, returns fallback
func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}

// GetInt(key, fallback int) int
// Looks for a key in Environment Variables
// If found, returns value as an int
// If not found, returns fallback
// If value is not an int, returns fallback
func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valAsInt
}
