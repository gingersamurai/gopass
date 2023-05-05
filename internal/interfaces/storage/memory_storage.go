package storage

import (
	"errors"
	"fmt"
	"gopass/internal/entity"
	"gopass/internal/usecase"
	"sync"
)

type MemoryStorage struct {
	sync.RWMutex
	cipher usecase.Cipher

	userData   map[int64]entity.User
	nextUserId int64

	accountData   map[int64]entity.Account
	nextAccountId int64
}

func NewMemoryStorage(cipher usecase.Cipher) *MemoryStorage {
	userData := make(map[int64]entity.User)
	accountData := make(map[int64]entity.Account)
	return &MemoryStorage{
		userData:    userData,
		accountData: accountData,
		cipher:      cipher,
	}
}

var (
	errNotFound    = errors.New("MemoryStorage: account not found")
	errCantDecrypt = errors.New("MemoryStorage: cant decrypt data")
	errCantEncrypt = errors.New("MemoryStorage: cant encrypt data")
)

func (ms *MemoryStorage) GetAccount(id int64) (entity.Account, error) {
	ms.RLock()
	defer ms.RUnlock()

	result, ok := ms.accountData[id]
	if !ok {
		return entity.Account{}, errNotFound
	}

	decryptedPassword, err := ms.cipher.Decrypt([]byte(result.Password))
	if err != nil {
		return entity.Account{}, fmt.Errorf("%w: %w", errCantDecrypt, err)
	}

	result.Password = string(decryptedPassword)

	return result, nil
}

func (ms *MemoryStorage) AddAccount(account entity.Account) (int64, error) {
	ms.Lock()
	defer ms.Unlock()

	encryptedPassword, err := ms.cipher.Encrypt([]byte(account.Password))
	if err != nil {
		return 0, fmt.Errorf("%w: %w", errCantEncrypt, err)
	}
	account.Password = string(encryptedPassword)

	account.Id = ms.nextAccountId
	ms.nextAccountId++
	ms.accountData[account.Id] = account
	return account.Id, nil
}

func (ms *MemoryStorage) DeleteAccount(id int64) error {
	if _, ok := ms.accountData[id]; ok {
		return errNotFound
	}

	delete(ms.accountData, id)
	return nil
}
