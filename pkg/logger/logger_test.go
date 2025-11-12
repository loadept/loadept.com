package logger

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/loadept/loadept.com/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	err := os.MkdirAll("logs", os.ModePerm)
	assert.NoError(t, err, "Should be able to create the directory records")

	defer os.RemoveAll("logs")

	os.Setenv("DEBUG", "false")
	defer os.Unsetenv("DEBUG")
	config.LoadEnviron()

	NewLogger()
	INFO.Println("Test message")
	ERROR.Println("Test error message")

	defer CloseLogger()

	currentDate := time.Now().Format("2006-01-02")
	infoPath := filepath.Join("logs", "access-"+currentDate+".log")
	errorPath := filepath.Join("logs", "error-"+currentDate+".log")

	_, err = os.Stat(infoPath)
	assert.NoError(t, err, "Info file should exist")

	_, err = os.Stat(errorPath)
	assert.NoError(t, err, "Error file should exist")

	infoContent, err := os.ReadFile(infoPath)
	assert.NoError(t, err)
	assert.Contains(t, string(infoContent), "Test message")

	errorContent, err := os.ReadFile(errorPath)
	assert.NoError(t, err)
	assert.Contains(t, string(errorContent), "Test error message")
}
