package memory_cache

import (
	"github.com/stretchr/testify/assert"
	"gopass/internal/entity"
	"log"
	"testing"
	"time"
)

func TestMemoryCache(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		mc := NewMemoryCache()
		myAuthData := entity.AuthData{
			UserId: 5,
			Key:    "bibaboba",
		}

		err := mc.AddKey(myAuthData.UserId, myAuthData.Key, 5*time.Second)
		assert.NoError(t, err)
		log.Println("done simpleadd")
		err = mc.AddKey(myAuthData.UserId, myAuthData.Key, 5*time.Second)
		assert.Error(t, err)

		key, err := mc.GetKey(myAuthData.UserId)
		assert.NoError(t, err)
		assert.Equal(t, myAuthData.Key, key)

		_, err = mc.GetKey(228)
		assert.Error(t, err)

		err = mc.DeleteKey(myAuthData.UserId)
		assert.NoError(t, err)

		_, err = mc.GetKey(myAuthData.UserId)
		assert.Error(t, err)

		myAuthData = entity.AuthData{
			UserId: 10,
			Key:    "yaSlavaMarlow",
		}
		lifetime := time.Second * 5
		err = mc.AddKey(myAuthData.UserId, myAuthData.Key, lifetime)
		assert.NoError(t, err)
		time.Sleep(lifetime + time.Second)
		_, err = mc.GetKey(myAuthData.UserId)
		assert.Error(t, err)
	})
}
