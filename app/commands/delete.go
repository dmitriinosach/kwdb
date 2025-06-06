package commands

import (
	"kwdb/app/storage"
)

const CommandDelete = "DELETE"

type DeleteCommand struct {
	name       string
	Args       *arguments
	isWritable bool
}

func NewDeleteCommand() CommandInterface {
	return &DeleteCommand{
		name:       CommandDelete,
		Args:       new(arguments),
		isWritable: true,
	}
}

func (c *DeleteCommand) CheckArgs() bool {
	if c.Args.Key == "" {
		return false
	}

	return true
}

func (c *DeleteCommand) Execute() ([]byte, error) {

	err := storage.Storage.Delete(c.Args.Key)
	if err != nil {
		return []byte{}, err
	}

	return []byte{}, nil
}

func (c *DeleteCommand) Name() string {
	return c.name
}

func (c *DeleteCommand) SetArgs(args *arguments) {
	c.Args = args
}
func (c *DeleteCommand) IsWritable() bool {
	return c.isWritable
}
