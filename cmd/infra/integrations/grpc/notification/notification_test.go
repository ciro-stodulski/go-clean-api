package notificationpbgrpc

import (
	"context"
	"go-clean-api/cmd/infra/integrations/grpc/notification/pb"
	mockgrpcuser "go-clean-api/cmd/shared/mocks/infra/integrations/grpc/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Notification_Verify(t *testing.T) {
	t.Run("should verify notification with succeffully", func(t *testing.T) {
		// make mock
		msg := "+msg de test+"
		mockPbGrpc := new(mockgrpcuser.MockService)
		ctx := context.Background()

		request := &pb.Request{
			Msg: msg,
		}

		response := &pb.Reponse{
			Event: &pb.Event{
				Name:     "name event",
				Describe: "describe event",
			},
		}

		mockPbGrpc.On("Verify", ctx, request).Return(response, nil)
		//

		// test func
		s := New(mockPbGrpc)
		err := s.Verify(msg)
		//

		// asserts
		assert.Nil(t, err)
		mockPbGrpc.AssertCalled(t, "Verify", ctx, request)
		//
	})
}
