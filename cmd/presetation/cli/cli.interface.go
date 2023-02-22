package cliinterface

type Command interface {
	Run(CliLine) error
	Err(error) error
	GetOptions() Options
}
