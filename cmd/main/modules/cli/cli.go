package climodule

import (
	cobradapter "go-api/cmd/infra/adapters/cobra"
	cliinterface "go-api/cmd/interface/cli"
	clilistusers "go-api/cmd/interface/cli/list-users"
	"go-api/cmd/main/container"
	"go-api/cmd/shared/env"
	"os"
)

type CliModule struct {
	Program cobradapter.CobraAdapter
}

func New() CliModule {
	cli := CliModule{
		Program: cobradapter.CobraAdapter{},
	}

	if os.Args[1] != "list-commands" {
		cli.Program.RunCommand(LoadCommands())
	}

	cli.Program.ListCommands(ListCommands())

	return cli
}

// for load command to execute
func LoadCommands() []cliinterface.Command {
	env.Load()

	c := container.New()

	return []cliinterface.Command{
		clilistusers.New(c.ListUsersUseCase),
	}
}

// for list commands in output
func ListCommands() []cliinterface.Command {
	return []cliinterface.Command{
		&clilistusers.ListUsersCli{},
	}
}
