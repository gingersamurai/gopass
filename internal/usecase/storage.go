package usecase

import (
	"errors"
	"gopass/internal/entity"
)

var (
	ErrServiceNotFound   = errors.New("storage: service not found")
	ErrAccountNotFound   = errors.New("storage: account not found")
	ErrUserNotFound      = errors.New("storage: user not found")
	ErrUserAlreadyExists = errors.New("storage: user with that login already exists")
	ErrWrongPassword     = errors.New("storage: wrong login or password")
)

type Storage interface {
	ServiceStorage
	UserStorage
	AccountStorage
}

type ServiceStorage interface {
	AddService(service entity.Service) (int64, error)
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
	AddAccount(key string, account entity.Account) (int64, error)
	GetAccount(key string, id int64) (entity.Account, error)
	DeleteAccount(id int64) error
}
