package commands

import (
	"kwdb/app/storage"
)

type DeleteCommand struct {
	name       string
	Args       CommandArguments
	isWritable bool
}

func (command *DeleteCommand) IsWritable() bool {
	return command.isWritable
}

func (command *DeleteCommand) CheckArgs(args CommandArguments) bool {
	if args.Key == "" {
		return false
	}

	return true
}

func (command *DeleteCommand) Execute() (string, error) {

	storage.Storage.DeleteValue(command.Args.Key)

	return "", nil
}

func (command *DeleteCommand) Name() string {
	return command.name
}

func (command *DeleteCommand) SetArgs(args CommandArguments) {
	command.Args = args
}
