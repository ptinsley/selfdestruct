package utils

import "os"

// EnvOr - return an ENV value OR a default value for fallback
func EnvOr(envVar string, defaultStr string) string {
	value := Env(envVar)
	if value == "" {
		value = defaultStr
	}

	return value
}

// Env - basic return of ENV, somewhat redundant
func Env(envVar string) string {
	return os.Getenv(envVar)
}
