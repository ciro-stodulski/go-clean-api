package verifynotificationusecase

import (
	portsservice "go-clean-api/cmd/core/ports"
	mockservicesnotification "go-clean-api/cmd/shared/mocks/infra/services/notification"
	"testing"
)

func Test_UseCase_Verify_Notification(t *testing.T) {
	t.Run("Notify with succeffully", func(t *testing.T) {
		// make mock services
		mockUserServices := new(mockservicesnotification.MockNotificationServices)

		dto := portsservice.Dto{
			Name:  "test",
			Event: "test",
		}

		mockUserServices.On("CheckNotify", dto.Name).Return(nil, nil)
		//

		// test func
		usecase := New(mockUserServices)
		usecase.Notify(dto)
		//

		// asserts
		mockUserServices.AssertCalled(t, "CheckNotify", dto.Name)
		//
	})
}
