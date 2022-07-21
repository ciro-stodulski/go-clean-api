package createuserusecase

import (
	create_dto "go-api/cmd/interface/amqp/consumers/users/create/dto"
	mocks "go-api/cmd/shared/mocks"
	mockservicesuser "go-api/cmd/shared/mocks/services/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UseCase_CreateUser(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		mockServices := new(mockservicesuser.MockServices)
		userMock := mocks.NewMockUser()

		mockServices.On("CreateUser").Return(userMock, nil)

		usecase := New(mockServices)

		dto := create_dto.CreateDto{
			Name:     "test",
			Email:    "test@test",
			Password: "test",
		}

		result, err := usecase.CreateUser(dto)

		assert.Nil(t, err)
		assert.Equal(t, userMock, result)
	})
}
