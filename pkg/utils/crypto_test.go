package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	password := "securepassword123"

	hash, err := HashPassword(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
	assert.NotEqual(t, password, hash)
}

func TestCheckPassword(t *testing.T) {
	password := "securepassword123"
	hash, _ := HashPassword(password)

	match, err := CheckPassword(password, hash)
	assert.NoError(t, err)
	assert.True(t, match)

	match, err = CheckPassword("wrongpassword", hash)
	assert.Equal(t, bcrypt.ErrMismatchedHashAndPassword, err)
	assert.False(t, match)
}
