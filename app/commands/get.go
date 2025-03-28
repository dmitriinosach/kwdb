package commands

import (
	"context"
	"kwdb/app/storage"
)

const CommandGet = "GET"

type GetCommand struct {
	name       string
	Args       *arguments
	isWritable bool
}

func NewGetCommand() *GetCommand {
	return &GetCommand{
		name:       CommandGet,
		Args:       new(arguments),
		isWritable: false,
	}
}

func (c *GetCommand) IsWritable(ctx context.Context) bool {
	return c.isWritable
}

func (c *GetCommand) SetArgs(ctx context.Context, args *arguments) {
	c.Args = args
}

func (c *GetCommand) CheckArgs(ctx context.Context, args *arguments) bool {
	if args.Key == "" {
		return false
	}

	return true
}

func (c *GetCommand) Name() string {

	return c.name
}

func (c *GetCommand) Execute(ctx context.Context) (string, error) {

	cell, err := storage.Storage.Get(ctx, c.Args.Key)

	if cell == nil {
		storage.Status.Metrics.Miss()

		return "", err
	}

	storage.Status.Metrics.Hit()

	return cell.Value, nil
}
