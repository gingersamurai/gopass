package encrypter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAESEncrypter(t *testing.T) {

	t.Run("small length of key", func(t *testing.T) {
		key := "12345678"
		aesc := NewAESEncrypter()
		_, err := aesc.Encrypt(key, "bibaboba")
		assert.Error(t, err)
	})

	t.Run("big length of key", func(t *testing.T) {
		key := "12345678901234567901234567890"
		aesc := NewAESEncrypter()
		_, err := aesc.Encrypt(key, "bibaboba")
		assert.Error(t, err)
	})

	t.Run("simple", func(t *testing.T) {
		key := "1234567890123456"
		data := "0123456789abcdef"

		myCipher := NewAESEncrypter()
		encrypted, err := myCipher.Encrypt(key, data)
		assert.NoError(t, err)
		decrypted, err := myCipher.Decrypt(key, encrypted)
		assert.Equal(t, data, decrypted)
	})

	t.Run("with padding", func(t *testing.T) {
		key := "1234567890123456"
		data := "0123456789abcdefbibaboba"

		myCipher := NewAESEncrypter()
		encrypted, err := myCipher.Encrypt(key, data)
		assert.NoError(t, err)
		decrypted, err := myCipher.Decrypt(key, encrypted)
		assert.Equal(t, data, decrypted)
	})

	t.Run("bad encoding", func(t *testing.T) {
		key := "1234567890123456"
		data := "бибабоба"

		myCipher := NewAESEncrypter()
		encrypted, err := myCipher.Encrypt(key, data)
		assert.NoError(t, err)
		decrypted, err := myCipher.Decrypt(key, encrypted)
		assert.Equal(t, data, decrypted)
	})

	t.Run("empty", func(t *testing.T) {
		key := "1234567890123456"
		data := ""

		myCipher := NewAESEncrypter()
		encrypted, err := myCipher.Encrypt(key, data)
		assert.NoError(t, err)
		decrypted, err := myCipher.Decrypt(key, encrypted)
		assert.Equal(t, data, decrypted)
	})
}
