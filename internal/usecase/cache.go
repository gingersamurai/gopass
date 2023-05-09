package usecase

type Cache interface {
	AuthDataCache
}

type AuthDataCache interface {
	SetKey(key string) error
	GetKey(key string) (string, error)
	DeleteKey(key string) error
}
