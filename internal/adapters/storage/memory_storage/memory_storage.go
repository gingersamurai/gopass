package memory_storage

import (
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
