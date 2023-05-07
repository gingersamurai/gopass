package memory_storage

import (
	"gopass/internal/entity"
	"gopass/internal/usecase"
	"sync"
)

type MemoryStorage struct {
	sync.RWMutex

	symmetricEncrypter  usecase.SymmetricEncrypter
	asymmetricEncrypter usecase.AsymmetricEncrypter

	serviceData map[int64]entity.Service
	userData    map[int64]entity.User
	accountData map[int64]entity.Account

	nextServiceId int64
	nextUserId    int64
	nextAccountId int64
}

func NewMemoryStorage(se usecase.SymmetricEncrypter, ae usecase.AsymmetricEncrypter) *MemoryStorage {
	serviceData := make(map[int64]entity.Service)
	userData := make(map[int64]entity.User)
	accountData := make(map[int64]entity.Account)
	return &MemoryStorage{
		serviceData:         serviceData,
		userData:            userData,
		accountData:         accountData,
		symmetricEncrypter:  se,
		asymmetricEncrypter: ae,
	}
}
