package commands

import (
	"context"
	"kwdb/app/storage"
)

const CommandDelete = "DELETE"

type DeleteCommand struct {
	name       string
	Args       *arguments
	isWritable bool
}

func NewDeleteCommand() *DeleteCommand {
	return &DeleteCommand{
		name:       CommandDelete,
		Args:       new(arguments),
		isWritable: true,
	}
}

func (c *DeleteCommand) CheckArgs(ctx context.Context, args *arguments) bool {
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

func (c *DeleteCommand) SetArgs(ctx context.Context, args *arguments) {
	c.Args = args
}
func (c *DeleteCommand) IsWritable(ctx context.Context) bool {
	return c.isWritable
}
