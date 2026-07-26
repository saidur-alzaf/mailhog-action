package main

import "os"

// envOrDefault returns the value of the environment variable named by key
// if it is set to a non-empty value. Otherwise, it returns fallback.
func envOrDefault(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}