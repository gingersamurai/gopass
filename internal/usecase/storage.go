package usecase

import "gopass/internal/entity"

type Storage interface {
	AccountStorage
	UserStorage
}

type AccountStorage interface {
	GetAccount(id int64) (entity.Account, error)
	AddAccount(account entity.Account) (int64, error)
	DeleteAccount(id int64) error
}

type UserStorage interface {
	GetUser(id int64) (entity.User, error)
	AddUser(user entity.User) (int64, error)
	DeleteUser(id int64) error
}
