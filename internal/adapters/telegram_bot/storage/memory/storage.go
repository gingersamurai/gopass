package memory

import (
	"fmt"
	"gopass/internal/adapters/telegram_bot/storage"
	"sync"
)

type TelegramUserStorage struct {
	sync.RWMutex
	data map[int64]int64
}

func NewTelegramUserStorage() *TelegramUserStorage {
	data := make(map[int64]int64)

	return &TelegramUserStorage{data: data}
}

func (ms *TelegramUserStorage) AddTelegramUser(tgUserId, userId int64) error {
	ms.Lock()
	defer ms.Unlock()

	if _, ok := ms.data[tgUserId]; ok {
		return fmt.Errorf("TelegramUserStorage.AddTelegramUser(): %w", storage.ErrTgUserAlreadyExists)
	}

	ms.data[tgUserId] = userId
	return nil
}

func (ms *TelegramUserStorage) GetTelegramUser(tgUserId int64) (int64, error) {
	ms.RLock()
	defer ms.RUnlock()

	if _, ok := ms.data[tgUserId]; !ok {
		return 0, fmt.Errorf("TelegramUserStorage.GetTelegramUser(): %w", storage.ErrTgUserNotFound)
	}

	return ms.data[tgUserId], nil
}

func (ms *TelegramUserStorage) DeleteTelegramUser(tgUserId int64) error {
	ms.Lock()
	defer ms.Unlock()

	if _, ok := ms.data[tgUserId]; !ok {
		return fmt.Errorf("TelegramUserStorage.DeleteTelegramUser(): %w", storage.ErrTgUserNotFound)
	}

	delete(ms.data, tgUserId)
	return nil
}
