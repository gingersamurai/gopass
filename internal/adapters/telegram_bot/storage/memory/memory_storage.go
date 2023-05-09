package memory

import (
	"fmt"
	"gopass/internal/adapters/telegram_bot/storage"
	"sync"
)

type MemoryTelegramUserStorage struct {
	sync.RWMutex
	data map[int64]int64
}

func (ms *MemoryTelegramUserStorage) AddTelegramUser(tgUserId, userId int64) error {
	ms.Lock()
	defer ms.Unlock()

	if _, ok := ms.data[tgUserId]; ok {
		return fmt.Errorf("MemoryTelegramUserStorage.AddTelegramUser(): %w", storage.ErrTgUserAlreadyExists)
	}

	ms.data[tgUserId] = userId
	return nil
}

func (ms *MemoryTelegramUserStorage) GetTelegramUser(tgUserId int64) (int64, error) {
	ms.RLock()
	defer ms.RUnlock()

	if _, ok := ms.data[tgUserId]; !ok {
		return 0, fmt.Errorf("MemoryTelegramUserStorage.GetTelegramUser(): %w", storage.ErrTgUserNotFound)
	}

	return ms.data[tgUserId], nil
}

func (ms *MemoryTelegramUserStorage) DeleteTelegramUser(tgUserId int64) error {
	ms.Lock()
	defer ms.Unlock()

	if _, ok := ms.data[tgUserId]; !ok {
		return fmt.Errorf("MemoryTelegramUserStorage.DeleteTelegramUser(): %w", storage.ErrTgUserNotFound)
	}

	delete(ms.data, tgUserId)
	return nil
}
