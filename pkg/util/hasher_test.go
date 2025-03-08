package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "mysecurepassword"

	hash, err := HashPassword(password)
	assert.NoError(t, err, "HasPassword should not return an error")
	assert.NotEmpty(t, hash, "Hash should not be empty")
	assert.NotEqual(t, password, hash, "Hash should not be equal to the plain password")

	_, err = HashPassword("")
	assert.Error(t, err, "HasPassword should return an error for empty password")
}

func TestCheckPasswordHash(t *testing.T) {
	password := "mysecurepassword"

	hash, err := HashPassword(password)
	assert.NoError(t, err, "HasPassword should not return an error")

	assert.True(t, CheckPasswordHash(password, hash), "CheckPasswordHash should return true for matching password and hash")

	assert.False(t, CheckPasswordHash("wrongpassword", hash), "CheckPasswordHash should return false for incorrect password")

	assert.False(t, CheckPasswordHash(password, ""), "CheckPasswordHash should return false for empty hash")
}
