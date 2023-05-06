package memory_storage

import "errors"

var (
	errAccountNotFound = errors.New("MemoryStorage: account not found")
	errUserNotFound    = errors.New("MemoryStorage: user not found")
	errCantEncrypt     = errors.New("MemoryStorage: cant encrypt data")
	errCantDecrypt     = errors.New("MemoryStorage: cant decrypt data")
)
