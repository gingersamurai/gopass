package encrypter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBcryptEncrypter(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		be := NewBcryptEncrypter()
		password := "bibaboba"
		hash, _ := be.Encrypt(password)
		result, err := be.Compare(hash, password)
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("wrong password", func(t *testing.T) {
		be := NewBcryptEncrypter()
		password := "bibaboba"
		hash, _ := be.Encrypt(password)
		wrongPassword := "AHAHAHAHAHAHAH NOT PANIC"
		result, err := be.Compare(hash, wrongPassword)
		assert.NoError(t, err)
		assert.False(t, result)
	})
}
