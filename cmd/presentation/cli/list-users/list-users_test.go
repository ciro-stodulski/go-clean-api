package clilistusers

import (
	"errors"
	cliinterface "go-clean-api/cmd/presentation/cli"
	usecasemock "go-clean-api/cmd/shared/mocks/application/use-case/use-case"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Cli_List_Users(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		// make mock
		mockUse := new(usecasemock.MockUseCase[any, any])

		mockUse.On("Perform", nil).Return(0, nil)
		//

		// test func
		cli := New(mockUse)
		err := cli.Run(cliinterface.CliLine{Line: ""})
		//

		// asserts
		assert.Nil(t, err)
		mockUse.AssertNumberOfCalls(t, "Perform", 1)
		//
	})

	t.Run("error INTERNAL_ERROR", func(t *testing.T) {
		// make mock
		mockUse := new(usecasemock.MockUseCase[any, any])
		//

		// test func
		cli := New(mockUse)

		err_internal := errors.New("internal errors")
		err := cli.Err(err_internal)
		//

		// asserts
		assert.NotNil(t, err)
		assert.Equal(t, err, err_internal)
		//
	})
}
