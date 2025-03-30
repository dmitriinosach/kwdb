package commands

import (
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

func (c *InfoCommand) IsWritable() bool {
	return c.isWritable
}

func (c *InfoCommand) SetArgs(args *arguments) {
	c.Args = args
}

func (c *InfoCommand) CheckArgs() bool {
	return true
}

func (c *InfoCommand) Name() string {

	return c.name
}

func (c *InfoCommand) Execute() (string, error) {
	return storage.Storage.Info(), nil
}
