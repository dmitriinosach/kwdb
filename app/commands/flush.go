package commands

import (
	"kwdb/app/storage"
)

const CommandFlush = "flush"

type FlushCommand struct {
	name       string
	Args       *arguments
	isWritable bool
}

func NewFlushCommand() CommandInterface {
	return &FlushCommand{
		name:       CommandStatus,
		isWritable: false,
	}
}

func (c *FlushCommand) IsWritable() bool {
	return c.isWritable
}

func (c *FlushCommand) SetArgs(args *arguments) {
	c.Args = args
}

func (c *FlushCommand) CheckArgs() bool {
	return true
}

func (c *FlushCommand) Name() string {

	return c.name
}

func (c *FlushCommand) Execute() ([]byte, error) {

	storage.Storage.Flush()

	return []byte{}, nil
}
