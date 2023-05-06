package usecase

type Cipher interface {
	Encrypt(message string) (string, error)
	Decrypt(messageCipher string) (string, error)
}
