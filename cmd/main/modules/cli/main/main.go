package main

import (
	climodule "go-clean-api/cmd/main/modules/cli"
)

func main() {
	cli := climodule.New()

	cli.Program.Execute()
}
