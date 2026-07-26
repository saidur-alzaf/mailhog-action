package main

import (
	"os"
	"testing"
)

func TestEnvOrDefault(t *testing.T) {
	t.Run("returns provided value when env var exists", func(t *testing.T) {
		t.Setenv("SMTP_FROM", "sender@example.com")

		if got := envOrDefault("SMTP_FROM", "fallback@example.com"); got != "sender@example.com" {
			t.Fatalf("expected env value, got %q", got)
		}
	})

	t.Run("returns default when env var is empty", func(t *testing.T) {
		t.Setenv("SMTP_FROM", "")

		if got := envOrDefault("SMTP_FROM", "fallback@example.com"); got != "fallback@example.com" {
			t.Fatalf("expected fallback value, got %q", got)
		}
	})

	t.Run("returns default when env var is unset", func(t *testing.T) {
		os.Unsetenv("SMTP_FROM")

		if got := envOrDefault("SMTP_FROM", "fallback@example.com"); got != "fallback@example.com" {
			t.Fatalf("expected fallback value, got %q", got)
		}
	})
}
