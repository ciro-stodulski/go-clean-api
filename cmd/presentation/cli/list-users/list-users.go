package clilistusers

import (
	domainusecases "go-clean-api/cmd/domain/use-case"
	cliinterface "go-clean-api/cmd/presentation/cli"
)

type ListUsersCli struct {
	luuc domainusecases.ListUsersUseCase
}

func (luc *ListUsersCli) GetOptions() cliinterface.Options {
	return cliinterface.Options{
		Command_name: "list-users",
		Description:  "command for list user",
	}
}

func New(luuc domainusecases.ListUsersUseCase) cliinterface.Command {

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
