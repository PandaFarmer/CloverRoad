package core

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHashAndSalt(t *testing.T) {
	user_password := "TopS3cr3tTopS3cr3tTopS3cr3t"

	hashedPwd, err := HashAndSalt(user_password)
	assert.NoError(t, err)
	os.Setenv("passwordHash", hashedPwd)
}

func TestComparePasswords(t *testing.T) {
	hash := os.Getenv("passwordHash")
	err := ComparePasswords(hash, "TopS3cr3tTopS3cr3tTopS3cr3t")
	assert.NoError(t, err)
}
