package encrypter

import "golang.org/x/crypto/bcrypt"

type BcryptEncrypter struct{}

func NewBcryptEncrypter() *BcryptEncrypter {
	return &BcryptEncrypter{}
}

func (be *BcryptEncrypter) Encrypt(message string) (string, error) {
	encryptedMessage, err := bcrypt.GenerateFromPassword([]byte(message), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encryptedMessage), nil
}

func (be *BcryptEncrypter) Compare(encryptedMessage string, message string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedMessage), []byte(message))
	if err != nil {
		return false, nil
	} else {
		return true, nil
	}
}
