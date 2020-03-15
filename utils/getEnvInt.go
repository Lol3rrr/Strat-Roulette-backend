package utils

import (
	"os"
	"strconv"
)

// GetEnvInt returns the Env-Variable for the key or
// the fallback if not Env-Variable is found
func GetEnvInt(key string, fallback int) int {
	rawValue, found := os.LookupEnv(key)
	if !found {
		return fallback
	}

	value, err := strconv.Atoi(rawValue)
	if err != nil {
		return fallback
	}

	return value
}
