package find_service

import (
	"context"
	"go-api/src/core/entities/user"
	"go-api/src/main/container"
	"go-api/src/presentation/grpc/services/user/pb"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserCase struct {
	mock.Mock
}

func (mock *MockUserCase) GetUser(id string) (*user.User, error) {
	arg := mock.Called()
	result := arg.Get(0)
	return result.(*user.User), arg.Error(1)
}

func newMockUser() *user.User {
	u, _ := user.New("test", "test", "test")
	return u
}

func Test_ServiceGrpc_FindUser_Create(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		mockRepo := new(MockUserCase)
		userMock := newMockUser()

		mockRepo.On("GetUser").Return(userMock, nil)

		testService := New(&container.Container{
			GetUserUseCase: mockRepo,
		})

		pb := &pb.NewRequestFindUser{ID: "tes"}
		ctx := context.Background()

		result, err := testService.FindUser(ctx, pb)

		assert.Nil(t, err)
		mockRepo.AssertCalled(t, "GetUser")
		assert.Equal(t, result.User.ID, userMock.ID.String())
		assert.Equal(t, result.User.Email, userMock.Email)
		assert.Equal(t, result.User.Name, userMock.Name)
		assert.Equal(t, result.User.CreatedAt, userMock.CreatedAt.String())
	})
}
