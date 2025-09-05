package commands

import (
	"kwdb/app/storage"
)

const CommandSet = "SET"

type SetCommand struct {
	name       string
	Args       *arguments
	isWritable bool
}

func NewSetCommand() CommandInterface {
	return &SetCommand{
		name:       CommandSet,
		Args:       new(arguments),
		isWritable: true,
	}
}

func (c *SetCommand) IsWritable() bool {
	return c.isWritable
}

func (c *SetCommand) CheckArgs() bool {
	if c.Args.Key == "" || c.Args.Value == nil {
		return false
	}
	return true
}

func (c *SetCommand) Execute() ([]byte, error) {
	err := storage.Storage.Set(c.Args.Key, c.Args.Value, c.Args.TTL)
	if err != nil {
		return []byte{}, err
	}

	return []byte("ok"), nil
}

func (c *SetCommand) Name() string {
	return c.name
}

func (c *SetCommand) SetArgs(args *arguments) {
	c.Args = args
}
