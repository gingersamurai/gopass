package usecase

import (
	"gopass/internal/entity"
)

type Storage interface {
	ServiceStorage
	UserStorage
	AccountStorage
}

type ServiceStorage interface {
	AddService(service entity.Service) (int64, error)
	GetService(id int64) (entity.Service, error)
	GetServiceByName(name string) (entity.Service, error)
	DeleteService(id int64) error
}

type UserStorage interface {
	AddUser(login, password string) (int64, error)
	GetUser(id int64) (entity.User, error)
	GetUserByLoginAndPassword(login, password string) (entity.User, error)
	DeleteUser(id int64) error
}

type AccountStorage interface {
	AddAccount(key string, account entity.Account) (int64, error)
	GetAccount(key string, id int64) (entity.Account, error)
	GetAccountsByServiceId(key string, serviceId int64) ([]entity.Account, error)
	DeleteAccount(id int64) error
}
