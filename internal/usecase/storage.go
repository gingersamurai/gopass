package usecase

import "gopass/internal/entity"

type Storage interface {
	ServiceStorage
	UserStorage
	AccountStorage
}

type ServiceStorage interface {
	AddService(user entity.Service) (int64, error)
	GetService(id int64) (entity.Service, error)
	DeleteService(id int64) error
}

type UserStorage interface {
	AddUser(login, password string) (int64, error)
	GetUser(id int64) (entity.User, error)
	GetUserByLoginAndPassword(login, password string) (entity.User, error)
	DeleteUser(id int64) error
}

type AccountStorage interface {
	AddAccount(account entity.Account) (int64, error)
	GetAccount(id int64) (entity.Account, error)
	DeleteAccount(id int64) error
}
