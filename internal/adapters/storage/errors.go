package storage

import "errors"

var (
	ErrServiceNotFound   = errors.New("storage: service not found")
	ErrAccountNotFound   = errors.New("storage: account not found")
	ErrUserNotFound      = errors.New("storage: user not found")
	ErrUserAlreadyExists = errors.New("storage: user with that login already exists")
	ErrWrongPassword     = errors.New("storage: wrong login or password")
)
