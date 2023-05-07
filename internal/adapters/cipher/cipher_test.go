package cipher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAESCipher(t *testing.T) {

	t.Run("small length of key", func(t *testing.T) {
		key := "12345678"
		_, err := NewAESCipher([]byte(key))
		assert.Error(t, err)
	})

	t.Run("big length of key", func(t *testing.T) {
		key := "12345678901234567901234567890"
		_, err := NewAESCipher([]byte(key))
		assert.Error(t, err)
	})

	t.Run("simple", func(t *testing.T) {
		key := "1234567890123456"
		data := "0123456789abcdef"
		myCipher, err := NewAESCipher([]byte(key))
		assert.NoError(t, err)

		encrypted, err := myCipher.Encrypt(data)
		assert.NoError(t, err)
		decrypted, err := myCipher.Decrypt(encrypted)
		assert.Equal(t, data, decrypted)
	})

	t.Run("with padding", func(t *testing.T) {
		key := "1234567890123456"
		data := "0123456789abcdefbibaboba"
		myCipher, err := NewAESCipher([]byte(key))
		assert.NoError(t, err)

		encrypted, err := myCipher.Encrypt(data)
		assert.NoError(t, err)
		decrypted, err := myCipher.Decrypt(encrypted)
		assert.Equal(t, data, decrypted)
	})

	t.Run("bad encoding", func(t *testing.T) {
		key := "1234567890123456"
		data := "бибабоба"
		myCipher, err := NewAESCipher([]byte(key))
		assert.NoError(t, err)

		encrypted, err := myCipher.Encrypt(data)
		assert.NoError(t, err)
		decrypted, err := myCipher.Decrypt(encrypted)
		assert.Equal(t, data, decrypted)
	})

	t.Run("empty", func(t *testing.T) {
		key := "1234567890123456"
		data := ""
		myCipher, err := NewAESCipher([]byte(key))
		assert.NoError(t, err)

		encrypted, err := myCipher.Encrypt(data)
		assert.NoError(t, err)
		decrypted, err := myCipher.Decrypt(encrypted)
		assert.Equal(t, data, decrypted)
	})
}
