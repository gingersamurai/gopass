package memory

import (
	"fmt"
	"gopass/internal/adapters/storage"
	"gopass/internal/entity"
)

func (ms *Storage) AddUser(login, password string) (int64, error) {
	ms.Lock()
	defer ms.Unlock()

	for _, user := range ms.userData {
		if user.Login == login {
			return 0, fmt.Errorf("memoryStorage.AddUser(): %w", storage.ErrUserAlreadyExists)
		}
	}

	encryptedPassword, err := ms.asymmetricEncrypter.Encrypt(password)
	if err != nil {
		return 0, fmt.Errorf("memoryStorage.AddUser(): %w", err)
	}

	user := entity.User{
		Id:           ms.nextUserId,
		Login:        login,
		PasswordHash: encryptedPassword,
	}
	ms.nextUserId++

	ms.userData[user.Id] = user

	return user.Id, nil
}

func (ms *Storage) GetUser(id int64) (entity.User, error) {
	ms.RLock()
	defer ms.RUnlock()

	user, ok := ms.userData[id]
	if !ok {
		return entity.User{}, fmt.Errorf("memoryStorage.GetUser(): %w", storage.ErrUserNotFound)
	}

	return user, nil
}

func (ms *Storage) GetUserByLoginAndPassword(login, password string) (entity.User, error) {
	ms.RLock()
	defer ms.RUnlock()

	for _, user := range ms.userData {
		goodPassword, err := ms.asymmetricEncrypter.Compare(user.PasswordHash, password)
		if err != nil {
			return entity.User{}, err
		}
		if user.Login == login && goodPassword {
			return user, nil
		}
	}

	return entity.User{}, fmt.Errorf(
		"memoryStorage.GetUserByLoginAndPassword(): %w",
		storage.ErrWrongPassword)
}

func (ms *Storage) DeleteUser(id int64) error {
	if _, ok := ms.userData[id]; !ok {
		return fmt.Errorf("memoryStorage.DeleteUser: %w", storage.ErrUserNotFound)
	}

	delete(ms.userData, id)
	return nil
}
