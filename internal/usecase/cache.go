package usecase

type Cache interface {
	AuthDataCache
}

type AuthDataCache interface {
	AddKey(userId int64, key string) error
	GetKey(userId int64) (string, error)
	DeleteKey(key string) error
}
