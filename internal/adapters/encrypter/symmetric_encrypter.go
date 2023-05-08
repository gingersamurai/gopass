package encrypter

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

type AESEncrypter struct{}

func NewAESEncrypter() *AESEncrypter {
	return &AESEncrypter{}
}

func (aesc *AESEncrypter) Encrypt(key, message string) (string, error) {

	paddedRawMessage := aesc.addPadding([]byte(message))

	myCipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("AESEncrypter.Encrypt(): %w", err)
	}
	cipherText := make([]byte, aes.BlockSize+len(paddedRawMessage))

	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("AESEncrypter.Encrypt(): %w", err)
	}

	mode := cipher.NewCBCEncrypter(myCipher, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], paddedRawMessage)

	return hex.EncodeToString(cipherText), nil
}

func (aesc *AESEncrypter) Decrypt(key, encryptedMessage string) (string, error) {
	rawMessageCipher, err := hex.DecodeString(encryptedMessage)
	if err != nil {
		return "", err
	}

	myCipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("AESEncrypter.Encrypt(): %w", err)
	}

	iv := rawMessageCipher[:aes.BlockSize]
	rawMessageCipher = rawMessageCipher[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(myCipher, iv)

	paddedRawMessage := make([]byte, len(rawMessageCipher))
	mode.CryptBlocks(paddedRawMessage, rawMessageCipher)
	rawMessage := aesc.removePadding(paddedRawMessage)

	return string(rawMessage), nil
}

func (aesc *AESEncrypter) addPadding(rawMessage []byte) []byte {
	paddingLength := aes.BlockSize - len(rawMessage)%aes.BlockSize
	padding := bytes.Repeat([]byte{byte(paddingLength)}, paddingLength)

	return append(rawMessage, padding...)
}

func (aesc *AESEncrypter) removePadding(paddedRawMessage []byte) []byte {
	padding := int(paddedRawMessage[len(paddedRawMessage)-1])
	return paddedRawMessage[:len(paddedRawMessage)-padding]
}
