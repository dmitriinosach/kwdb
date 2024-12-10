package commands

import (
	"kwdb/app/storage"
)

type InfoCommand struct {
	name string
	Args CommandArguments
}

func (command *InfoCommand) SetArgs(args CommandArguments) {
	command.Args = args
}

func (command *InfoCommand) CheckArgs(args CommandArguments) bool {
	return true
}

func (command *InfoCommand) Name() string {

	return command.name
}

func (command *InfoCommand) Execute() (string, error) {

	storage.Info()

	return "--", nil
}
