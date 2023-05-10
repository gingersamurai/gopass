package memory

import (
	"github.com/stretchr/testify/assert"
	"gopass/internal/adapters/encrypter"
	"gopass/internal/adapters/storage/memory"
	"gopass/internal/usecase"
	"testing"
)

func TestTelegramUserStorage(t *testing.T) {
	ae := encrypter.NewBcryptEncrypter()
	se := encrypter.NewAESEncrypter()
	ms := memory.New(se, ae)
	uin := usecase.NewUserInteractor(ms, ae)

	ebumbaData := struct {
		tgId  int64
		login string
	}{
		tgId:  5,
		login: "ginger.ebumba",
	}
	id, _, err := uin.RegisterUser(ebumbaData.login)
	assert.NoError(t, err)

	mtus := NewTelegramUserStorage()
	err = mtus.AddTelegramUser(ebumbaData.tgId, id)
	assert.NoError(t, err)

	err = mtus.AddTelegramUser(ebumbaData.tgId, id)
	assert.Error(t, err)

	gotId, err := mtus.GetTelegramUser(ebumbaData.tgId)
	assert.NoError(t, err)
	assert.Equal(t, gotId, id)

	_, err = mtus.GetTelegramUser(3452)
	assert.Error(t, err)

	err = mtus.DeleteTelegramUser(ebumbaData.tgId)
	assert.NoError(t, err)

	_, err = mtus.GetTelegramUser(ebumbaData.tgId)
	assert.Error(t, err)

	err = mtus.DeleteTelegramUser(ebumbaData.tgId)
	assert.Error(t, err)
}
