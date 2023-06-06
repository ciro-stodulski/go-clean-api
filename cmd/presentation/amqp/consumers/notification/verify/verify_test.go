package verifyconsumer

import (
	domaindto "go-clean-api/cmd/domain/dto"
	"go-clean-api/cmd/domain/exception"
	"go-clean-api/cmd/main/container"
	verifynotificationusecasemock "go-clean-api/cmd/shared/mocks/application/use-cases/verify-notification"

	ports_amqp "go-clean-api/cmd/presentation/amqp/ports"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Consumer_verify(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		// make mock
		mockUseCase := new(verifynotificationusecasemock.MockUseCase)

		dto := domaindto.Event{
			Name:  "test",
			Event: "test",
		}

		mockUseCase.On("Notify", dto).Return((*exception.ApplicationException)(nil), nil)
		//

		// test func
		testService := New(&container.Container{
			VerifyUseCase: mockUseCase,
		})
		//

		err := testService.MessageHandler(ports_amqp.Message{
			Body: dto,
		})

		// asserts
		assert.Nil(t, err)
		mockUseCase.AssertCalled(t, "Notify", dto)
		//
	})
}
