package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashing(t *testing.T) {
	password := "mySecretPin"
	hash, err := HashString(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
	assert.True(t, CheckStringHash(password, hash))
}
