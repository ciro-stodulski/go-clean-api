package clilistusers

import (
	listusersusecase "go-api/cmd/core/use-case/list-user"
	cliinterface "go-api/cmd/interface/cli"
)

type listUsersCli struct {
	luuc listusersusecase.ListUsersUseCase
}

func (luc *listUsersCli) GetOptions() cliinterface.Options {
	return cliinterface.Options{
		Command_name: "list-users",
		Description:  "command for list user",
	}
}

func New(luuc listusersusecase.ListUsersUseCase) cliinterface.Command {

	return &listUsersCli{
		luuc: luuc,
	}
}

func (luc *listUsersCli) Run(line cliinterface.CliLine) error {
	luc.luuc.ListUsers()

	return nil
}

func (luc *listUsersCli) Err(err error) error {

	return err
}
