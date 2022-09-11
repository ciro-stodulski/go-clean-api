package clilistusers

import (
	"errors"
	cliinterface "go-api/cmd/interface/cli"
	listuserusecasemock "go-api/cmd/shared/mocks/core/use-cases/list-user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Cli_List_Users(t *testing.T) {
	t.Run("succeffully", func(t *testing.T) {
		// make mock
		mockUse := new(listuserusecasemock.MockUseCase)

		mockUse.On("ListUsers").Return(nil, nil)
		//

		// test func
		cli := New(mockUse)
		err := cli.Run(cliinterface.CliLine{Line: ""})
		//

		// asserts
		assert.Nil(t, err)
		mockUse.AssertNumberOfCalls(t, "ListUsers", 1)
		//
	})

	t.Run("error INTERNAL_ERROR", func(t *testing.T) {
		// make mock
		mockUse := new(listuserusecasemock.MockUseCase)
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