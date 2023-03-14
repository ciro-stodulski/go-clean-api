package userservice

import (
	user "go-clean-api/cmd/domain/entities/user"
	mockhttpjsonplaceholder "go-clean-api/cmd/shared/mocks/infra/integrations/http/jsonplaceholder"
	mockusercache "go-clean-api/cmd/shared/mocks/infra/repositories/cache/user"
	mocksqluser "go-clean-api/cmd/shared/mocks/infra/repositories/sql/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Service_Register(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		mr := new(mocksqluser.MockRepository)
		mi := new(mockhttpjsonplaceholder.MockIntegration)
		mockCache := new(mockusercache.MockCache)

		mr.On("Create").Return(nil)

		service := New(mr, mi, mockCache)

		u := &user.User{
			Name:     "test",
			Email:    "test@test",
			Password: "test",
		}

		result, errApp, err := service.Register(u)

		assert.Nil(t, err)
		assert.Nil(t, errApp)
		assert.Equal(t, u.Name, result.Name)
		assert.Equal(t, u.Email, result.Email)
	})
}
