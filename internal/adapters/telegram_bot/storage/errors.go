package storage

import "errors"

var (
	ErrTgUserAlreadyExists = errors.New("telegram user already exists")
	ErrTgUserNotFound      = errors.New("telegram user not found")
)
