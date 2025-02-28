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
		isWritable: true,
	}
}

func (c *SetCommand) IsWritable(ctx context.Context) bool {
	return c.isWritable
}

func (c *SetCommand) CheckArgs(ctx context.Context, args *CommandArguments) bool {
	if args.Key == "" || args.Value == "" {
		return false
	}

	return true
}

func (c *SetCommand) Execute(ctx context.Context) (string, error) {

	err := storage.Storage.Set(ctx, c.Args.Key, c.Args.Value, c.Args.TTL)
	if err != nil {
		return "", err
	}

	return "OK", nil
}

func (c *SetCommand) Name() string {
	return c.name
}

func (c *SetCommand) SetArgs(ctx context.Context, args *CommandArguments) {
	c.Args = args
}
