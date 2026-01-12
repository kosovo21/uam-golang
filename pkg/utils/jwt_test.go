package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	email := "test@example.com"
	userID := "123"
	secret := "supersecretkey"
	expiration := time.Minute * 15

	token, err := GenerateToken(email, userID, secret, expiration)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestValidateToken(t *testing.T) {
	email := "test@example.com"
	userID := "123"
	secret := "supersecretkey"
	expiration := time.Minute * 15

	token, _ := GenerateToken(email, userID, secret, expiration)

	claims, err := ValidateToken(token, secret)

	assert.NoError(t, err)
	assert.NotNil(t, claims)
	assert.Equal(t, email, claims.Email)
	assert.Equal(t, userID, claims.UserID)
}

func TestValidateToken_Expired(t *testing.T) {
	email := "test@example.com"
	userID := "123"
	secret := "supersecretkey"
	expiration := -time.Minute // Expired

	token, _ := GenerateToken(email, userID, secret, expiration)

	claims, err := ValidateToken(token, secret)

	assert.Error(t, err)
	assert.Nil(t, claims)
	assert.Contains(t, err.Error(), "token is expired")
}

func TestValidateToken_InvalidSecret(t *testing.T) {
	email := "test@example.com"
	userID := "123"
	secret := "supersecretkey"
	wrongSecret := "wrongsecret"
	expiration := time.Minute * 15

	token, _ := GenerateToken(email, userID, secret, expiration)

	claims, err := ValidateToken(token, wrongSecret)

	assert.Error(t, err)
	assert.Nil(t, claims)
	assert.ErrorContains(t, err, "signature is invalid")
}
