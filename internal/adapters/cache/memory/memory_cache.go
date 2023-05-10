package memory

import (
	"fmt"
	"gopass/internal/adapters/cache"
	"gopass/internal/entity"
	"sync"
	"time"
)

type Cache struct {
	sync.RWMutex
	data map[int64]entity.AuthData
}

func New() *Cache {
	data := make(map[int64]entity.AuthData)
	return &Cache{data: data}
}

func (mc *Cache) GetKey(userId int64) (string, error) {
	mc.RLock()
	defer mc.RUnlock()

	if authData, ok := mc.data[userId]; !ok {
		return "", fmt.Errorf("memoryCache.GetKey(): %w", cache.ErrNotFound)
	} else {
		return authData.Key, nil
	}
}

func (mc *Cache) AddKey(userId int64, key string, lifetime time.Duration) error {
	mc.Lock()
	defer mc.Unlock()

	if _, ok := mc.data[userId]; ok {
		return fmt.Errorf("memoryCache.AddKey(): %w", cache.ErrAlreadyExists)
	}

	mc.data[userId] = entity.AuthData{
		UserId: userId,
		Key:    key,
	}

	go func() {
		time.Sleep(lifetime)
		_ = mc.DeleteKey(userId)
	}()
	return nil
}

func (mc *Cache) DeleteKey(userId int64) error {
	mc.Lock()
	defer mc.Unlock()

	if _, ok := mc.data[userId]; !ok {
		return fmt.Errorf("memoryCache.AddKey(): %w", cache.ErrNotFound)
	}

	delete(mc.data, userId)
	return nil
}
