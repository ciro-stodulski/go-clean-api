package registeruserusecase

import (
	domaindto "go-clean-api/cmd/domain/dto"
	entity "go-clean-api/cmd/domain/entities"
	portsservice "go-clean-api/cmd/domain/services"
	mocks "go-clean-api/cmd/shared/mocks"
	mockservicesnotification "go-clean-api/cmd/shared/mocks/infra/services/notification"
	mockservicesuser "go-clean-api/cmd/shared/mocks/infra/services/user"

	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func Test_UseCase_RegisterUser(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		// stub methods time.Now(), entity.NewID() and bcrypt.GenerateFromPassword()
		datapatch := monkey.Patch(time.Now,
			func() time.Time { return time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC) })

		uuidpatch := monkey.Patch(entity.NewID,
			func() uuid.UUID {
				return uuid.MustParse("af687f58-5421-4a1f-aae5-6869a0d768f2")
			})

		userpatch := monkey.Patch(bcrypt.GenerateFromPassword,
			func([]byte, int) ([]byte, error) {
				return []byte("af687f58-5421-4a1f-aae5-6869a0d768f2"), nil
			})

		defer datapatch.Unpatch()
		defer uuidpatch.Unpatch()
		defer userpatch.Unpatch()
		//

		// make mock services
		mockUserServices := new(mockservicesuser.MockUserServices)
		mockNotificationServices := new(mockservicesnotification.MockNotificationServices)

		dto := domaindto.Dto{
			Name:     "test",
			Email:    "test@test",
			Password: "test",
		}
		user_mock := mocks.CreateMockUser(dto.Email, dto.Password, dto.Name)

		notificationFake := portsservice.Dto{Name: "REGISTERED_USER", Event: "USER"}
		idFake := "63494fdabb1e0bf59fb8fc5b"

		mockUserServices.On("Register", user_mock).Return(user_mock, nil)
		mockNotificationServices.On("SendNotify", notificationFake).Return(nil)
		mockNotificationServices.On("SaveNotify", notificationFake).Return("63494fdabb1e0bf59fb8fc5b")
		mockNotificationServices.On("FindById", idFake).Return(&notificationFake)
		//

		// test func
		usecase := New(mockUserServices, mockNotificationServices)
		result, err := usecase.Register(dto)
		//

		// asserts
		assert.Nil(t, err)
		assert.Equal(t, user_mock, result)
		mockUserServices.AssertCalled(t, "Register", user_mock)
		mockNotificationServices.AssertCalled(t, "SendNotify", notificationFake)
		mockNotificationServices.AssertCalled(t, "SaveNotify", notificationFake)
		mockNotificationServices.AssertCalled(t, "FindById", idFake)
		//
	})
}