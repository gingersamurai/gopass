package usecase

import (
	"github.com/stretchr/testify/assert"
	"gopass/internal/adapters/encrypter"
	"gopass/internal/adapters/storage/memory"
	"testing"
)

func TestUserInteractor(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		se := encrypter.NewAESEncrypter()
		ae := encrypter.NewBcryptEncrypter()

		storage := memory.NewMemoryStorage(se, ae)

		uin := NewUserInteractor(storage, ae)

		login := "ebumba"

		id, key, err := uin.RegisterUser(login)
		assert.NoError(t, err)
		user, err := uin.storage.GetUser(id)

		needServiceName := "telegram_bot"
		needLogin := "Ebumba_e"
		needPassword := "bibaboba123"
		id, err = uin.Set(user, key, needServiceName, needLogin, needPassword)
		assert.NoError(t, err)

		service, err := uin.storage.GetServiceByName(needServiceName)
		assert.NoError(t, err)
		needServiceId := service.Id

		accounts, err := uin.Get(user, key, needServiceName)
		assert.NoError(t, err)
		assert.Len(t, accounts, 1)
		assert.Equal(t, accounts[0].ServiceId, needServiceId)
		assert.Equal(t, accounts[0].Login, needLogin)
		assert.Equal(t, accounts[0].Password, needPassword)

		err = uin.Del(user, key, accounts[0].Id)
		assert.NoError(t, err)

		accounts, err = uin.Get(user, key, needServiceName)
		assert.Error(t, err)
	})
}
