package memory_storage

import (
	"gopass/internal/entity"
)

func (ms *MemoryStorage) GetUser(id int64) (entity.User, error) {
	ms.RLock()
	defer ms.RUnlock()

	result, ok := ms.userData[id]
	if !ok {
		return entity.User{}, errUserNotFound
	}

	decryptedPassword, err := ms.cipher.Decrypt([]byte(result.Password))
	if err != nil {
		return entity.User{}, errCantDecrypt
	}

	result.Password = string(decryptedPassword)

	return result, nil
}

func (ms *MemoryStorage) AddUser(user entity.User) (int64, error) {
	ms.Lock()
	defer ms.Unlock()

	encryptedPassword, err := ms.cipher.Encrypt([]byte(user.Password))
	if err != nil {
		return 0, errCantEncrypt
	}

	user.Password = string(encryptedPassword)

	user.Id = ms.nextUserId
	ms.nextUserId++
	ms.userData[user.Id] = user

	return user.Id, nil
}

func (ms *MemoryStorage) DeleteUser(id int64) error {
	if _, ok := ms.userData[id]; !ok {
		return errUserNotFound
	}

	delete(ms.userData, id)
	return nil
}
