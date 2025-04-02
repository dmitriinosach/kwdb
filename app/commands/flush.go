package commands

const CommandFlush = "flush"

type FlushCommand struct {
	name       string
	Args       *arguments
	isWritable bool
}

func NewFlushCommand() *FlushCommand {
	return &FlushCommand{
		name:       CommandStatus,
		Args:       new(arguments),
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

func (c *FlushCommand) Execute() (string, error) {

	return "1", nil
}
