package verifyconsumer

import (
	portsservice "go-api/cmd/core/ports"
	"go-api/cmd/main/container"
	verifynotificationusecasemock "go-api/cmd/shared/mocks/core/use-cases/verify-notification"

	ports_amqp "go-api/cmd/interface/amqp/ports"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Consumer_verify(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		// make mock
		mockUseCase := new(verifynotificationusecasemock.MockUseCase)

		dto := portsservice.Dto{
			Name:  "test",
			Event: "test",
		}

		mockUseCase.On("Notify", dto).Return(nil)
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
