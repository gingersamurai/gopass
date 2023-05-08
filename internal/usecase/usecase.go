package usecase

import (
	"errors"
	"fmt"
	"gopass/internal/entity"
)

var (
	ErrWrongPassword = errors.New("wrong password")
)

type UserInteractor struct {
	storage Storage
	ae      AsymmetricEncrypter
}

func NewUserInteractor(storage Storage, ae AsymmetricEncrypter) *UserInteractor {
	return &UserInteractor{
		storage: storage,
		ae:      ae}
}

func (uin *UserInteractor) RegisterUser(login string) (int64, string, error) {
	key, err := GenerateKey()
	if err != nil {
		return 0, "", fmt.Errorf("userInteractor.RegisterUser(): %w", err)
	}
	id, err := uin.storage.AddUser(login, key)
	if err != nil {
		return 0, "", fmt.Errorf("userInteractor.RegisterUser(): %w", err)
	}
	return id, key, nil
}

func (uin *UserInteractor) CheckUser(user entity.User, key string) bool {
	result, _ := uin.ae.Compare(user.PasswordHash, key)
	return result
}

func (uin *UserInteractor) Set(user entity.User, key, serviceName, login, password string) (int64, error) {
	if !uin.CheckUser(user, key) {
		return 0, ErrWrongPassword
	}

	service, err := uin.storage.GetServiceByName(serviceName)
	if err != nil {
		id, err := uin.storage.AddService(entity.Service{Name: serviceName})
		if err != nil {
			return 0, fmt.Errorf("userInteractor.Set(): %w", err)
		}
		service, err = uin.storage.GetService(id)
		if err != nil {
			return 0, fmt.Errorf("userInteractor.Set(): %w", err)
		}
	}

	id, err := uin.storage.AddAccount(key, entity.Account{
		UserId:    user.Id,
		ServiceId: service.Id,
		Login:     login,
		Password:  password,
	})

	if err != nil {
		return 0, fmt.Errorf("userInteractor.Set(): %w", err)
	}

	return id, nil
}

func (uin *UserInteractor) Get(user entity.User, key, serviceName string) ([]entity.Account, error) {
	if !uin.CheckUser(user, key) {
		return nil, ErrWrongPassword
	}

	service, err := uin.storage.GetServiceByName(serviceName)
	if err != nil {
		return nil, fmt.Errorf("userInteractor.Get(): %w", err)
	}

	accounts, err := uin.storage.GetAccountsByServiceId(key, service.Id)
	if err != nil {
		return nil, fmt.Errorf("userInteractor.Get(): %w", err)
	}
	if len(accounts) == 0 {
		return nil, errors.New("userInteractor.Get(): accounts not found")
	}
	return accounts, nil
}

func (uin *UserInteractor) Del(user entity.User, key string, accountId int64) error {
	if !uin.CheckUser(user, key) {
		return ErrWrongPassword
	}

	err := uin.storage.DeleteAccount(accountId)
	if err != nil {
		return fmt.Errorf("userInteractor.Del(): %w", err)
	}
	return nil
}
