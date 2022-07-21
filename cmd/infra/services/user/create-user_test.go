package userservice

import (
	user "go-api/cmd/core/entities/user"
	create_dto "go-api/cmd/interface/amqp/consumers/users/create/dto"
	mocks "go-api/cmd/shared/mocks"
	mockhttpjsonplaceholder "go-api/cmd/shared/mocks/integrations/http/jsonplaceholder"
	mockusercache "go-api/cmd/shared/mocks/repositories/cache/user"
	mocksqluser "go-api/cmd/shared/mocks/repositories/sql/user"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_Service_CreateUser(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		mr := new(mocksqluser.MockRepository)
		mi := new(mockhttpjsonplaceholder.MockIntegration)
		umock := &user.User{ID: uuid.Nil}
		mockCache := new(mockusercache.MockCache)

		mr.On("GetByEmail").Return(umock, nil)
		mr.On("Create")

		usecase := New(mr, mi, mockCache)

		dto := create_dto.CreateDto{
			Name:     "test",
			Email:    "test@test",
			Password: "test",
		}

		result, err := usecase.CreateUser(dto)

		assert.Nil(t, err)
		assert.Equal(t, dto.Name, result.Name)
		assert.Equal(t, dto.Email, result.Email)
	})

	t.Run("user already exists", func(t *testing.T) {
		mr := new(mocksqluser.MockRepository)
		mi := new(mockhttpjsonplaceholder.MockIntegration)
		userMockResult := mocks.NewMockUser()
		mockCache := new(mockusercache.MockCache)

		mr.On("GetByEmail").Return(userMockResult, nil)
		mr.On("Create")

		usecase := New(mr, mi, mockCache)

		dto := create_dto.CreateDto{
			Name:     "test",
			Email:    "test@test",
			Password: "test",
		}

		result, err := usecase.CreateUser(dto)

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), user.ErrUserAlreadyExists.Error())
	})
}
