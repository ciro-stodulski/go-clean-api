package userservice

import (
	user "go-api/cmd/core/entities/user"
	mocks "go-api/cmd/shared/mocks"
	mockhttpjsonplaceholder "go-api/cmd/shared/mocks/integrations/http/jsonplaceholder"
	mockusercache "go-api/cmd/shared/mocks/repositories/cache/user"
	mocksqluser "go-api/cmd/shared/mocks/repositories/sql/user"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_Service_Register(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		mr := new(mocksqluser.MockRepository)
		mi := new(mockhttpjsonplaceholder.MockIntegration)
		umock := &user.User{ID: uuid.Nil}
		mockCache := new(mockusercache.MockCache)

		mr.On("GetByEmail").Return(umock, nil)
		mr.On("Create")

		service := New(mr, mi, mockCache)

		u := &user.User{
			Name:     "test",
			Email:    "test@test",
			Password: "test",
		}

		result, err := service.Register(u)

		assert.Nil(t, err)
		assert.Equal(t, u.Name, result.Name)
		assert.Equal(t, u.Email, result.Email)
	})

	t.Run("user already exists", func(t *testing.T) {
		mr := new(mocksqluser.MockRepository)
		mi := new(mockhttpjsonplaceholder.MockIntegration)
		userMockResult := mocks.NewMockUser()
		mockCache := new(mockusercache.MockCache)

		mr.On("GetByEmail").Return(userMockResult, nil)
		mr.On("Create")

		service := New(mr, mi, mockCache)

		u := &user.User{
			Name:     "test",
			Email:    "test@test",
			Password: "test",
		}
		result, err := service.Register(u)

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), user.ErrUserAlreadyExists.Error())
	})
}
