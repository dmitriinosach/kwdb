package commands

import (
	"context"
	"kwdb/app/storage"
)

const CommandDelete = "DELETE"

type DeleteCommand struct {
	name       string
	Args       *Arguments
	isWritable bool
}

func NewDeleteCommand() *DeleteCommand {
	return &DeleteCommand{
		name:       CommandDelete,
		Args:       new(Arguments),
		isWritable: true,
	}
}

func (c *DeleteCommand) CheckArgs(ctx context.Context, args *Arguments) bool {
	if args.Key == "" {
		return false
	}

	return true
}

func (c *DeleteCommand) Execute(ctx context.Context) (string, error) {

	err := storage.Storage.Delete(ctx, c.Args.Key)
	if err != nil {
		return "", err
	}

	return "", nil
}

func (c *DeleteCommand) Name() string {
	return c.name
}

func (c *DeleteCommand) SetArgs(ctx context.Context, args *Arguments) {
	c.Args = args
}
func (c *DeleteCommand) IsWritable(ctx context.Context) bool {
	return c.isWritable
}
