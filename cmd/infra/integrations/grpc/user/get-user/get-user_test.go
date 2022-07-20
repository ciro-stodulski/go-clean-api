package getuserservice

import (
	"context"
	"go-api/cmd/infra/integrations/grpc/user/get-user/pb"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type MockService struct {
	mock.Mock
}

func (mock *MockService) GetUser(c context.Context, pbm *pb.NewRequestGetUser, op ...grpc.CallOption) (*pb.NewResponseGetUser, error) {
	arg := mock.Called(c, pbm, "test")
	result := arg.Get(0)
	return result.(*pb.NewResponseGetUser), arg.Error(1)
}

func newMockUsers() *pb.NewResponseGetUser {
	return &pb.NewResponseGetUser{
		Customer: &pb.Customer{
			ID:        "346",
			Name:      "test",
			Email:     "test@test",
			CreatedAt: "2022-05-01",
		},
	}
}

func Test_GetUserService_GetUser(t *testing.T) {
	t.Run("should get user with succeffully", func(t *testing.T) {
		userMock := newMockUsers()
		mockService := new(MockService)
		ctx := context.Background()

		request := &pb.NewRequestGetUser{
			ID: userMock.Customer.ID,
		}

		mockService.On("GetUser", ctx, request, "test").Return(userMock, nil)

		s := New(mockService)

		result, err := s.GetUser(userMock.Customer.ID)

		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
}
