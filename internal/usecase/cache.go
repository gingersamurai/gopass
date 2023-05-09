package usecase

import "time"

type Cache interface {
	AuthDataCache
}

type AuthDataCache interface {
	AddKey(userId int64, key string, lifetime time.Duration) error
	GetKey(userId int64) (string, error)
	DeleteKey(userId int64) error
}
