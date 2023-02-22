package climodule

import (
	cobradapter "go-clean-api/cmd/infra/adapters/cobra"
	"go-clean-api/cmd/main/container"
	cliinterface "go-clean-api/cmd/presetation/cli"
	clilistusers "go-clean-api/cmd/presetation/cli/list-users"
	"go-clean-api/cmd/shared/env"
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
