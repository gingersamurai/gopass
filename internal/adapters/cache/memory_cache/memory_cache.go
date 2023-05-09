package memory_cache

import (
	"fmt"
	"gopass/internal/adapters/cache"
	"gopass/internal/entity"
	"sync"
	"time"
)

type MemoryCache struct {
	sync.RWMutex
	data map[int64]entity.AuthData
}

func NewMemoryCache() *MemoryCache {
	data := make(map[int64]entity.AuthData)
	return &MemoryCache{data: data}
}

func (mc *MemoryCache) GetKey(userId int64) (string, error) {
	mc.RLock()
	defer mc.RUnlock()

	if authData, ok := mc.data[userId]; !ok {
		return "", fmt.Errorf("memoryCache.GetKey(): %w", cache.ErrNotFound)
	} else {
		return authData.Key, nil
	}
}

func (mc *MemoryCache) AddKey(userId int64, key string) error {
	mc.Lock()
	defer mc.Unlock()

	if _, err := mc.GetKey(userId); err == nil {
		return fmt.Errorf("memoryCache.AddKey(): %w", cache.ErrAlreadyExists)
	}

	mc.data[userId] = entity.AuthData{
		UserId: userId,
		Key:    key,
	}
	go func() {
		time.Sleep(time.Second * 60 * 5)
		_ = mc.DeleteKey(userId)
	}()
	return nil
}

func (mc *MemoryCache) DeleteKey(userId int64) error {
	mc.Lock()
	defer mc.Unlock()

	if _, err := mc.GetKey(userId); err != nil {
		return fmt.Errorf("memoryCache.AddKey(): %w", cache.ErrNotFound)
	}

	delete(mc.data, userId)
}
