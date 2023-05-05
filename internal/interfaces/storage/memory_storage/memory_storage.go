package memory_storage

import (
	"errors"
	"gopass/internal/entity"
	"sync"
)

type MemoryStorage struct {
	sync.RWMutex

	userData   map[int64]entity.User
	nextUserId int64

	accountData   map[int64]entity.Account
	nextAccountId int64
}

func NewMemoryStorage() *MemoryStorage {
	userData := make(map[int64]entity.User)
	accountData := make(map[int64]entity.Account)
	return &MemoryStorage{
		userData:    userData,
		accountData: accountData,
	}
}

var errNotFound = errors.New("MemoryStorage: account not found")

func (ms *MemoryStorage) GetAccount(id int64) (entity.Account, error) {
	ms.RLock()
	defer ms.RUnlock()

	result, ok := ms.accountData[id]
	if !ok {
		return entity.Account{}, errNotFound
	}

	return result, nil
}

func (ms *MemoryStorage) AddAccount(account entity.Account) (int64, error) {
	ms.Lock()
	defer ms.Unlock()

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
