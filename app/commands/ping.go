package commands

const CommandPing = "ping"

type PingCommand struct {
	name string
}

func NewPingCommand() CommandInterface {
	return &PingCommand{
		name: CommandPing,
	}
}

func (c *PingCommand) IsWritable() bool {
	return false
}

func (c *PingCommand) SetArgs(args *arguments) {
	return
}

func (c *PingCommand) CheckArgs() bool {
	return true
}

func (c *PingCommand) Name() string {

	return c.name
}

func (c *PingCommand) Execute() ([]byte, error) {
	return []byte("pong"), nil
}
