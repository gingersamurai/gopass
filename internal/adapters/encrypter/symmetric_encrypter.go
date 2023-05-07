package encrypter

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

var (
	errCannotCreate = errors.New("aes encrypter: cannot create encrypter")
)

type AESEncrypter struct {
	cipher cipher.Block
}

func NewAESEncrypter(key string) (*AESEncrypter, error) {
	result, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, fmt.Errorf("%w: %w", errCannotCreate, err)
	}

	return &AESEncrypter{
		cipher: result,
	}, nil
}

func (aesc *AESEncrypter) addPadding(rawMessage []byte) []byte {
	paddingLength := aesc.cipher.BlockSize() - len(rawMessage)%aesc.cipher.BlockSize()
	padding := bytes.Repeat([]byte{byte(paddingLength)}, paddingLength)

	return append(rawMessage, padding...)
}

func (aesc *AESEncrypter) removePadding(paddedRawMessage []byte) []byte {
	padding := int(paddedRawMessage[len(paddedRawMessage)-1])
	return paddedRawMessage[:len(paddedRawMessage)-padding]
}

func (aesc *AESEncrypter) Encrypt(message string) (string, error) {

	paddedRawMessage := aesc.addPadding([]byte(message))

	cipherText := make([]byte, aesc.cipher.BlockSize()+len(paddedRawMessage))

	iv := cipherText[:aesc.cipher.BlockSize()]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("AESEncrypter.Encrypt(): %w", err)
	}

	mode := cipher.NewCBCEncrypter(aesc.cipher, iv)
	mode.CryptBlocks(cipherText[aesc.cipher.BlockSize():], paddedRawMessage)

	return hex.EncodeToString(cipherText), nil
}

func (aesc *AESEncrypter) Decrypt(encryptedMessage string) (string, error) {
	rawMessageCipher, err := hex.DecodeString(encryptedMessage)
	if err != nil {
		return "", err
	}

	iv := rawMessageCipher[:aesc.cipher.BlockSize()]
	rawMessageCipher = rawMessageCipher[aesc.cipher.BlockSize():]

	mode := cipher.NewCBCDecrypter(aesc.cipher, iv)

	paddedRawMessage := make([]byte, len(rawMessageCipher))
	mode.CryptBlocks(paddedRawMessage, rawMessageCipher)
	rawMessage := aesc.removePadding(paddedRawMessage)

	return string(rawMessage), nil
}
