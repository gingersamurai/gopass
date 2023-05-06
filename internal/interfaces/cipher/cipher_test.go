package cipher

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestAESCipher(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		key := "1234567890123456"
		data := "secret bibaboba abacaba"
		myCipher, err := NewAESCipher([]byte(key))
		assert.NoError(t, err)

		encrypted, err := myCipher.Encrypt(data)
		assert.NoError(t, err)
		log.Println("encrypted:", encrypted)
		decrypted, err := myCipher.Decrypt(encrypted)
		log.Println("decrypted:", decrypted)
		assert.Equal(t, data, decrypted)
	})
}
