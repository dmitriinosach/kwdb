package commands

import "kwdb/app/storage"

const CommandTruncate = "TRUNCATE"

type TruncateCommand struct {
	name       string
	Args       *arguments
	isWritable bool
}

func NewTruncateCommand() CommandInterface {
	return &TruncateCommand{
		name:       CommandTruncate,
		Args:       new(arguments),
		isWritable: false,
	}
}
func (c *TruncateCommand) CheckArgs() bool {
	return true
}

func (c *TruncateCommand) Execute() ([]byte, error) {
	storage.Storage.Truncate()

	return []byte{}, nil
}

func (c *TruncateCommand) Name() string {
	return c.name
}

func (c *TruncateCommand) SetArgs(args *arguments) {
	c.Args = args
}

func (c *TruncateCommand) IsWritable() bool {
	return c.isWritable
}
