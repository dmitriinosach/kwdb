package commands

import "kwdb/app/storage"

const CommandTruncate = "TRUNCATE"

type TruncateCommand struct {
	name       string
	Args       *arguments
	isWritable bool
}

func NewTruncateCommand() *TruncateCommand {
	return &TruncateCommand{
		name:       CommandTruncate,
		Args:       new(arguments),
		isWritable: false,
	}
}
func (c *TruncateCommand) CheckArgs() bool {
	return true
}

func (c *TruncateCommand) Execute() (string, error) {
	storage.Storage.Truncate()

	return "ok", nil
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
