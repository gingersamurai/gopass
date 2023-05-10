package memory

import (
	"gopass/internal/adapters/storage"
	"gopass/internal/entity"
	"sync"
)

type Storage struct {
	sync.RWMutex

	symmetricEncrypter  storage.SymmetricEncrypter
	asymmetricEncrypter storage.AsymmetricEncrypter

	serviceData map[int64]entity.Service
	userData    map[int64]entity.User
	accountData map[int64]entity.Account

	nextServiceId int64
	nextUserId    int64
	nextAccountId int64
}

func New(se storage.SymmetricEncrypter, ae storage.AsymmetricEncrypter) *Storage {
	serviceData := make(map[int64]entity.Service)
	userData := make(map[int64]entity.User)
	accountData := make(map[int64]entity.Account)
	return &Storage{
		serviceData:         serviceData,
		userData:            userData,
		accountData:         accountData,
		symmetricEncrypter:  se,
		asymmetricEncrypter: ae,
	}
}
