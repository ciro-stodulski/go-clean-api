package cliinterface

type CliLine struct {
	Line string
}

type Options struct {
	Command_name string
	Description  string
	Schema       any
}
