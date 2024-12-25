package commands

import (
	"fmt"
	"kwdb/app/storage"
)

type GetCommand struct {
	name       string
	Args       CommandArguments
	isWritable bool
}

func (command *GetCommand) IsWritable() bool {
	return command.isWritable
}

func (command *GetCommand) SetArgs(args CommandArguments) {
	command.Args = args
}

func (command *GetCommand) CheckArgs(args CommandArguments) bool {
	if args.Key == "" {
		return false
	}

	return true
}

func (command *GetCommand) Name() string {

	return command.name
}

func (command *GetCommand) Execute() (string, error) {

	if !storage.Storage.HasKey(command.Args.Key) {
		return "", fmt.Errorf("key %s does not exist", command.Args.Key)
	}

	value, _ := storage.Storage.GetValue(command.Args.Key)

	return value.Value, nil
}
