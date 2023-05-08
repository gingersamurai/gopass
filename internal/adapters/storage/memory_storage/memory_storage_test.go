package memory_storage

import (
	"github.com/stretchr/testify/assert"
	"gopass/internal/adapters/encrypter"
	"gopass/internal/entity"
	"testing"
)

func TestMemoryStorage(t *testing.T) {
	t.Run("simple user", func(t *testing.T) {
		login := "gingersamurai"
		secretKey := "0123456789abcdef"

		asynEnc := encrypter.NewBcryptEncrypter()

		synEnc := encrypter.NewAESEncrypter()

		ms := NewMemoryStorage(synEnc, asynEnc)

		id, err := ms.AddUser(login, secretKey)
		assert.NoError(t, err)

		user, err := ms.GetUserByLoginAndPassword(login, secretKey)
		assert.NoError(t, err)
		ok, err := asynEnc.Compare(user.PasswordHash, secretKey)
		assert.NoError(t, err)
		assert.True(t, ok)
		assert.Equal(t, user.Id, id)

		userById, err := ms.GetUser(id)
		assert.NoError(t, err)
		assert.Equal(t, userById, user)

		err = ms.DeleteUser(id)
		assert.NoError(t, err)
		_, err = ms.GetUserByLoginAndPassword(login, secretKey)
		assert.Error(t, err)
		_, err = ms.GetUser(id)
		assert.Error(t, err)
	})

	t.Run("simple service", func(t *testing.T) {

		asynEnc := encrypter.NewBcryptEncrypter()

		synEnc := encrypter.NewAESEncrypter()

		ms := NewMemoryStorage(synEnc, asynEnc)

		myService := entity.Service{Name: "tiktok"}
		id, err := ms.AddService(myService)
		assert.NoError(t, err)

		service, err := ms.GetService(id)
		assert.NoError(t, err)
		assert.Equal(t, service, myService)

		err = ms.DeleteService(id)
		assert.NoError(t, err)

		_, err = ms.GetService(id)
		assert.Error(t, err)
	})

	t.Run("simple account", func(t *testing.T) {
		secretKey := "0123456789abcdef"

		asynEnc := encrypter.NewBcryptEncrypter()

		synEnc := encrypter.NewAESEncrypter()
		ms := NewMemoryStorage(synEnc, asynEnc)

		myAccount := entity.Account{UserId: 1, ServiceId: 2, Login: "ebumba", Password: "bibaboba123"}
		id, err := ms.AddAccount(secretKey, myAccount)
		assert.NoError(t, err)

		account, err := ms.GetAccount(secretKey, id)
		assert.NoError(t, err)
		assert.Equal(t, myAccount, account)

		err = ms.DeleteAccount(id)
		assert.NoError(t, err)

		_, err = ms.GetAccount(secretKey, id)
		assert.Error(t, err)
	})
}
