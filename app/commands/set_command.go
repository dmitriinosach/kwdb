package commands

import (
	"kwdb/app/storage"
)

type SetCommand struct {
	name string
	Args CommandArguments
}

func (command *SetCommand) CheckArgs(args CommandArguments) bool {
	if args.Key == "" || args.Value == "" {
		return false
	}

	return true
}

func (command *SetCommand) Execute() (string, error) {

	storage.SetValue(command.Args.Key, command.Args.Value, 0)

	return "", nil
}

func (command *SetCommand) Name() string {
	return command.name
}

func (command *SetCommand) SetArgs(args CommandArguments) {
	command.Args = args
}
