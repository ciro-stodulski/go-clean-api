package notificaitonpb

import (
	"context"
	"go-clean-api/cmd/presentation/grpc/notification/pb"
	"go-clean-api/cmd/shared/mocks"
	usecasemock "go-clean-api/cmd/shared/mocks/application/use-case"

	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ServiceGrpc_FindUser_Create(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		// make mock
		mockUse := new(usecasemock.MockUseCase[any, any])
		userMock := mocks.NewMockUser()

		mockUse.On("Perform", nil).Return(userMock, nil)
		//

		// test func
		testPb := New(mockUse)

		pb := &pb.ResquestNotification{List: &pb.List{Name: "", Describe: ""}}
		ctx := context.Background()

		result, err := testPb.Verify(ctx, pb)
		//

		// asserts
		assert.Nil(t, err)
		mockUse.AssertCalled(t, "Perform", nil)
		assert.NotNil(t, result)
		//
	})
}
