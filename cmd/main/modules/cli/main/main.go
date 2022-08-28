package main

import (
	climodule "go-api/cmd/main/modules/cli"
)

func main() {
	cli := climodule.New()

	cli.Program.Cli.Execute()
}
