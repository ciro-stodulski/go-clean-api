package getuserservice

import (
	"context"
	"go-api/cmd/infra/integrations/grpc/user/get-user/pb"
	mockgrpcuser "go-api/cmd/shared/mocks/infra/integrations/grpc/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetUserService_GetUser(t *testing.T) {
	t.Run("should get user with succeffully", func(t *testing.T) {
		userMock := &pb.NewResponseGetUser{
			Customer: &pb.Customer{
				ID:        "346",
				Name:      "test",
				Email:     "test@test",
				CreatedAt: "2022-05-01",
			},
		}
		mockService := new(mockgrpcuser.MockService)
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
