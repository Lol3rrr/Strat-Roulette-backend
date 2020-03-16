package utils

import "os"

// GetEnvString returns the Env-Variable for the key or
// the fallback if not Env-Variable is found
func GetEnvString(key, fallback string) string {
	value, found := os.LookupEnv(key)
	if !found {
		return fallback
	}

	return value
}
