package commands

import (
	"context"
	"kwdb/app/storage"
)

const CommandDelete = "DELETE"

type DeleteCommand struct {
	name       string
	Args       *CommandArguments
	isWritable bool
}

func NewDeleteCommand() *GetCommand {
	return &GetCommand{
		name:       CommandDelete,
		Args:       new(CommandArguments),
		isWritable: false,
	}
}

func (command *DeleteCommand) IsWritable(ctx context.Context) bool {
	return command.isWritable
}

func (command *DeleteCommand) CheckArgs(ctx context.Context, args *CommandArguments) bool {
	if args.Key == "" {
		return false
	}

	return true
}

func (command *DeleteCommand) Execute(ctx context.Context) (string, error) {

	storage.Storage.DeleteValue(command.Args.Key)

	return "", nil
}

func (command *DeleteCommand) Name() string {
	return command.name
}

func (command *DeleteCommand) SetArgs(ctx context.Context, args *CommandArguments) {
	command.Args = args
}
