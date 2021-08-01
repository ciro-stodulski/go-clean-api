package v1_user

import (
	"go-api/src/core/entities/user"
	user_use_case "go-api/src/core/useCases/user"
	"go-api/src/main/container"
	ports_http "go-api/src/presentation/http/ports"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserServiceMock struct {
	mock.Mock
}

func (mock *UserServiceMock) GetById(id string) (*user.User, error) {
	arg := mock.Called()
	result := arg.Get(0)
	return result.(*user.User), arg.Error(1)
}

func newMockUser() *user.User {
	user, _ := user.NewUser("test", "test", "test")
	return user
}
func Test_GetUser(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		userMock := newMockUser()
		mockRepo := new(UserServiceMock)

		mockRepo.On("GetUser").Return(userMock, nil)

		container := &container.Container{
			UserService: user_use_case.Service{},
		}

		testService := NewController(container)

		id := "752ea551-5e6a-4382-859c-cd09fbe50110"
		result, err := testService.LoadRoute().Handle(ports_http.HttpRequest{
			Params: ports_http.Params{
				ports_http.Param{Key: "id", Value: id},
			},
		})

		assert.Nil(t, err)
		assert.NotNil(t, result)

	})

	// 	t.Run("error internal", func(t *testing.T) {

	// 	})

	// 	t.Run("error user not found", func(t *testing.T) {

	// 	})
}
