package mockgrpcuser

import (
	"context"
	"go-api/cmd/infra/integrations/grpc/user/get-user/pb"

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
