package clilistusers

import (
	usecase "go-clean-api/cmd/domain/use-case"
	cliinterface "go-clean-api/cmd/presentation/cli"
)

type ListUsersCli struct {
	listUsersUseCase usecase.UseCase[any, any]
}

func (luc *ListUsersCli) GetOptions() cliinterface.Options {
	return cliinterface.Options{
		Command_name: "list-users",
		Description:  "command for list user",
	}
}

func New(listUsersUseCase usecase.UseCase[any, any]) cliinterface.Command {

	return &ListUsersCli{
		listUsersUseCase,
	}
}

func (luc *ListUsersCli) Run(line cliinterface.CliLine) error {
	luc.listUsersUseCase.Perform(nil)

	return nil
}

func (luc *ListUsersCli) Err(err error) error {

	return err
}
