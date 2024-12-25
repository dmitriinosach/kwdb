package commands

import (
	"context"
	"kwdb/app/storage"
)

const CommandSet = "SET"

type SetCommand struct {
	name       string
	Args       *CommandArguments
	isWritable bool
}

func NewSetCommand() *SetCommand {
	return &SetCommand{
		name:       CommandSet,
		Args:       new(CommandArguments),
		isWritable: false,
	}
}

func (command *SetCommand) IsWritable(ctx context.Context) bool {
	return command.isWritable
}

func (command *SetCommand) CheckArgs(ctx context.Context, args *CommandArguments) bool {
	if args.Key == "" || args.Value == "" {
		return false
	}

	return true
}

func (command *SetCommand) Execute(ctx context.Context) (string, error) {

	storage.Storage.SetValue(command.Args.Key, command.Args.Value, command.Args.TTL)

	return "", nil
}

func (command *SetCommand) Name() string {
	return command.name
}

func (command *SetCommand) SetArgs(ctx context.Context, args *CommandArguments) {
	command.Args = args
}
