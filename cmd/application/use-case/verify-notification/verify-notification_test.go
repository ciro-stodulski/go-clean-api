package verifynotificationusecase

import (
	domaindto "go-clean-api/cmd/domain/dto"
	"go-clean-api/cmd/domain/exception"
	mockservicesnotification "go-clean-api/cmd/shared/mocks/infra/service/notification"
	"testing"
)

func Test_UseCase_Verify_Notification(t *testing.T) {
	t.Run("Notify with succeffully", func(t *testing.T) {
		// make mock services
		mockUserServices := new(mockservicesnotification.MockNotificationServices)

		dto := domaindto.Event{
			Name:  "test",
			Event: "test",
		}

		mockUserServices.On("CheckNotify", dto.Name).Return((*exception.ApplicationException)(nil), nil)
		//

		// test func
		usecase := New(mockUserServices)
		usecase.Perform(dto)
		//

		// asserts
		mockUserServices.AssertCalled(t, "CheckNotify", dto.Name)
		//
	})
}
