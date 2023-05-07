package usecase

type SymmetricEncrypter interface {
	Encrypt(message string) (string, error)
	Decrypt(encryptedMessage string) (string, error)
}

type AsymmetricEncrypter interface {
	Encrypt(message string) (string, error)
	Compare(encryptedMessage string, message string) (bool, error)
}