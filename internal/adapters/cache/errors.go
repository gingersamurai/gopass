package cache

import "errors"

var (
	ErrNotFound      = errors.New("authData not found")
	ErrAlreadyExists = errors.New("authData already exists")
)
