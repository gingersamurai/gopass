package storage

type SymmetricEncrypter interface {
	Encrypt(key, message string) (string, error)
	Decrypt(key, encryptedMessage string) (string, error)
}

type AsymmetricEncrypter interface {
	Encrypt(message string) (string, error)
	Compare(encryptedMessage string, message string) (bool, error)
}
