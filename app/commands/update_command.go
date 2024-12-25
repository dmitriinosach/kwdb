package commands

import (
	"kwdb/app/storage"
)

type UpdateCommand struct {
	name       string
	Args       CommandArguments
	isWritable bool
}

func (command *UpdateCommand) IsWritable() bool {
	return command.isWritable
}

func (command *UpdateCommand) CheckArgs(args CommandArguments) bool {
	if args.Key == "" || args.Value == "" {
		return false
	}

	return true
}

func (command *UpdateCommand) Execute() (string, error) {

	cell, _ := storage.Storage.GetValue(command.Args.Key)
	if cell != nil {
		cell.Value = command.Args.Value
	}

	return "", nil
}

func (command *UpdateCommand) Name() string {
	return command.name
}

func (command *UpdateCommand) SetArgs(args CommandArguments) {
	command.Args = args
}
