package v1_delete_user

import (
	"go-api/src/main/container"
	ports_http "go-api/src/presentation/http/ports"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserCase struct {
	mock.Mock
}

func (mock *MockUserCase) DeleteUser(id string) error {
	arg := mock.Called()
	return arg.Error(0)
}

func Test_Controller_Delete_Create(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		mockRepo := new(MockUserCase)
		id := "752ea551-5e6a-4382-859c-cd09fbe50110"

		mockRepo.On("DeleteUser").Return(nil)

		testService := NewController(&container.Container{
			DeleteUserUseCase: mockRepo,
		})

		result, err := testService.LoadRoute().Handle(ports_http.HttpRequest{
			Params: ports_http.Params{
				ports_http.Param{Key: "id", Value: id},
			},
		})

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, &ports_http.HttpResponse{
			Status: 204,
		}, result)
	})
}
