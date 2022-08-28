package clilistusers

import (
	listusersusecase "go-api/cmd/core/use-case/list-user"
	cliinterface "go-api/cmd/interface/cli"
)

type ListUsersCli struct {
	luuc listusersusecase.ListUsersUseCase
}

func (luc *ListUsersCli) GetOptions() cliinterface.Options {
	return cliinterface.Options{
		Command_name: "list-users",
		Description:  "command for list user",
	}
}

func New(luuc listusersusecase.ListUsersUseCase) cliinterface.Command {

	return &ListUsersCli{
		luuc: luuc,
	}
}

func (luc *ListUsersCli) Run(line cliinterface.CliLine) error {
	luc.luuc.ListUsers()

	return nil
}

func (luc *ListUsersCli) Err(err error) error {

	return err
}
