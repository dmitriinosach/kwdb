package commands

import (
	"context"
	"kwdb/app/storage"
)

const CommandInfo = "INFO"

type InfoCommand struct {
	name       string
	Args       *arguments
	isWritable bool
}

func NewInfoCommand() *InfoCommand {
	return &InfoCommand{
		name:       CommandInfo,
		Args:       new(arguments),
		isWritable: false,
	}
}

func (c *InfoCommand) IsWritable(ctx context.Context) bool {
	return c.isWritable
}

func (c *InfoCommand) SetArgs(ctx context.Context, args *arguments) {
	c.Args = args
}

func (c *InfoCommand) CheckArgs(ctx context.Context, args *arguments) bool {
	return true
}

func (c *InfoCommand) Name() string {

	return c.name
}

func (c *InfoCommand) Execute(ctx context.Context) (string, error) {
	return storage.Storage.Info(), nil
}
