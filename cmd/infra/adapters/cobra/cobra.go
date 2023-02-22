package cobradapter

import (
	"encoding/json"
	cliinterface "go-clean-api/cmd/presentation/cli"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type CobraAdapter struct {
	Cli cobra.Command
}

func (ca *CobraAdapter) ListCommands(cos []cliinterface.Command) {
	lc := &cobra.Command{
		Use:   "list-commands",
		Short: "List commands already to use",
		Run: func(cmd *cobra.Command, _ []string) {
			log.Default().Println("# Available commands to execute with 'run-command':")
			for _, command := range cos {
				log.Default().Println(formatListCommands(command.GetOptions()))
			}
		},
	}

	ca.Cli.AddCommand(lc)
}

func formatListCommands(options cliinterface.Options) string {
	format_command :=
		"| command:" + options.Command_name + " | description:" + options.Description

	if options.Schema != nil {
		out, _ := json.Marshal(options.Schema)

		format_command += " | schema:" + string(out)
	}

	format_command += "\n"

	return format_command
}

func (ca *CobraAdapter) RunCommand(cos []cliinterface.Command) {
	rc := &cobra.Command{
		Use:   "run-command",
		Short: "run-command <command>",
		Run: func(cmd *cobra.Command, args []string) {
			for _, command := range cos {
				if command.GetOptions().Command_name == args[0] {
					line := ""

					if len(args) > 1 {
						line = args[1]
					}

					err := command.Run(cliinterface.CliLine{
						Line: line,
					})

					if err != nil {
						err = command.Err(err)
						log.Default().Println(err.Error())
						os.Exit(1)
					}
				}
			}
		},
	}

	ca.Cli.AddCommand(rc)
}
