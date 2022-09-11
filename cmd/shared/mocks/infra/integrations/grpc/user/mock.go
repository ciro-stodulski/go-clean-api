package mockgrpcuser

import (
	"context"
	"go-clean-api/cmd/infra/integrations/grpc/notification/pb"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type MockService struct {
	mock.Mock
}

func (mock *MockService) Verify(c context.Context, pbm *pb.Request, op ...grpc.CallOption) (*pb.Reponse, error) {
	arg := mock.Called(c, pbm)
	result := arg.Get(0)
	return result.(*pb.Reponse), arg.Error(1)
}
