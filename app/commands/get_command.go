package commands

import (
	"errors"
	"kwdb/app/storage"
)

type GetCommand struct {
	name string
	Args CommandArguments
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

	if !storage.HasKey(command.Args.Key) {
		return "", errors.New("key not found")
	}

	value := storage.GetValue(command.Args.Key)

	return value.Value, nil
}
