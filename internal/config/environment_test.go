package config

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	os.Setenv("TEST_KEY", "some-value")
	defer os.Unsetenv("TEST_KEY")

	if got := getEnv("TEST_KEY", "default"); got != "some-value" {
		t.Errorf("getEnv() = %s; want some-value", got)
	}

	if got := getEnv("NON_EXISTENT", "defaultValue"); got != "defaultValue" {
		t.Errorf("getEnv() = %s; want defaultValue", got)
	}
}
