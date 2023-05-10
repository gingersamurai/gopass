package memory

import (
	"fmt"
	"gopass/internal/adapters/storage"
	"gopass/internal/entity"
)

func (ms *Storage) AddAccount(key string, account entity.Account) (int64, error) {
	ms.Lock()
	defer ms.Unlock()

	encryptedPassword, err := ms.symmetricEncrypter.Encrypt(key, account.Password)
	if err != nil {
		return 0, fmt.Errorf("memoryStorage.AddAccount(): %w", err)
	}
	account.Password = encryptedPassword

	account.Id = ms.nextAccountId
	ms.nextAccountId++
	ms.accountData[account.Id] = account
	return account.Id, nil
}

func (ms *Storage) GetAccount(key string, id int64) (entity.Account, error) {
	ms.RLock()
	defer ms.RUnlock()

	result, ok := ms.accountData[id]
	if !ok {
		return entity.Account{}, fmt.Errorf("memoryStorage.GetAccount(): %w", storage.ErrAccountNotFound)
	}

	decryptedPassword, err := ms.symmetricEncrypter.Decrypt(key, result.Password)
	if err != nil {
		return entity.Account{}, fmt.Errorf("memoryStorage.GetAccount(): %w", err)
	}

	result.Password = decryptedPassword

	return result, nil
}

func (ms *Storage) GetAccountsByServiceId(key string, serviceId int64) ([]entity.Account, error) {
	ms.RLock()
	defer ms.RUnlock()

	var result []entity.Account
	for _, account := range ms.accountData {
		if account.ServiceId != serviceId {
			continue
		}

		decryptedPassword, err := ms.symmetricEncrypter.Decrypt(key, account.Password)
		if err != nil {
			return nil, fmt.Errorf("memoryStorage.GetAccountsByServiceId(): %w", err)
		}
		account.Password = decryptedPassword

		result = append(result, account)
	}

	return result, nil
}

func (ms *Storage) DeleteAccount(id int64) error {
	ms.Lock()
	defer ms.Unlock()

	if _, ok := ms.accountData[id]; !ok {
		return fmt.Errorf("memoryStorage.DeleteAccount(): %w", storage.ErrAccountNotFound)
	}

	delete(ms.accountData, id)
	return nil
}
