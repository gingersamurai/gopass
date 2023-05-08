package memory_storage

import (
	"fmt"
	"gopass/internal/entity"
	"gopass/internal/usecase"
)

func (ms *MemoryStorage) AddAccount(account entity.Account) (int64, error) {
	ms.Lock()
	defer ms.Unlock()

	encryptedPassword, err := ms.symmetricEncrypter.Encrypt(account.Password)
	if err != nil {
		return 0, fmt.Errorf("memoryStorage.AddAccount(): %w", err)
	}
	account.Password = encryptedPassword

	account.Id = ms.nextAccountId
	ms.nextAccountId++
	ms.accountData[account.Id] = account
	return account.Id, nil
}

func (ms *MemoryStorage) GetAccount(id int64) (entity.Account, error) {
	ms.RLock()
	defer ms.RUnlock()

	result, ok := ms.accountData[id]
	if !ok {
		return entity.Account{}, fmt.Errorf("memoryStorage.GetAccount(): %w", usecase.ErrAccountNotFound)
	}

	decryptedPassword, err := ms.symmetricEncrypter.Decrypt(result.Password)
	if err != nil {
		return entity.Account{}, fmt.Errorf("memoryStorage.GetAccount(): %w", err)
	}

	result.Password = decryptedPassword

	return result, nil
}

func (ms *MemoryStorage) DeleteAccount(id int64) error {
	ms.Lock()
	defer ms.Unlock()

	if _, ok := ms.accountData[id]; ok {
		return fmt.Errorf("memoryStorage.DeleteAccount(): %w", usecase.ErrAccountNotFound)
	}

	delete(ms.accountData, id)
	return nil
}
