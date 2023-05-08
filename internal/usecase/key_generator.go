package usecase

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func GenerateKey() (string, error) {
	result := make([]byte, 12)
	_, err := io.ReadFull(rand.Reader, result)
	if err != nil {
		return "", fmt.Errorf("generateKey(): %w", err)
	}
	return base64.StdEncoding.EncodeToString(result), nil
}
