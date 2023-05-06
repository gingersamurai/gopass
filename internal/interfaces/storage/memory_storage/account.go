package memory_storage

import (
	"fmt"
	"gopass/internal/entity"
)

func (ms *MemoryStorage) GetAccount(id int64) (entity.Account, error) {
	ms.RLock()
	defer ms.RUnlock()

	result, ok := ms.accountData[id]
	if !ok {
		return entity.Account{}, errAccountNotFound
	}

	decryptedPassword, err := ms.cipher.Decrypt([]byte(result.Password))
	if err != nil {
		return entity.Account{}, fmt.Errorf("%w: %w", errCantDecrypt, err)
	}

	result.Password = string(decryptedPassword)

	return result, nil
}

func (ms *MemoryStorage) AddAccount(account entity.Account) (int64, error) {
	ms.Lock()
	defer ms.Unlock()

	encryptedPassword, err := ms.cipher.Encrypt([]byte(account.Password))
	if err != nil {
		return 0, fmt.Errorf("%w: %w", errCantEncrypt, err)
	}
	account.Password = string(encryptedPassword)

	account.Id = ms.nextAccountId
	ms.nextAccountId++
	ms.accountData[account.Id] = account
	return account.Id, nil
}

func (ms *MemoryStorage) DeleteAccount(id int64) error {
	ms.Lock()
	defer ms.Unlock()

	if _, ok := ms.accountData[id]; ok {
		return errAccountNotFound
	}

	delete(ms.accountData, id)
	return nil
}
