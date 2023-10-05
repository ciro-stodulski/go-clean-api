package verifyconsumer

import (
	"go-clean-api/cmd/domain/dto"
	domaindto "go-clean-api/cmd/domain/dto"
	"go-clean-api/cmd/domain/exception"
	usecasemock "go-clean-api/cmd/shared/mocks/application/use-case/use-case"

	ports_amqp "go-clean-api/cmd/presentation/amqp/ports"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Consumer_verify(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		// make mock
		mockUse := new(usecasemock.MockUseCase[dto.Event, any])

		dto := domaindto.Event{
			Name:  "test",
			Event: "test",
		}

		mockUse.On("Perform", dto).Return((*exception.ApplicationException)(nil), nil)
		//

		// test func
		testService := New(mockUse)
		//

		err := testService.MessageHandler(ports_amqp.Message{
			Body: dto,
		})

		// asserts
		assert.Nil(t, err)
		mockUse.AssertCalled(t, "Perform", dto)
		//
	})
}
