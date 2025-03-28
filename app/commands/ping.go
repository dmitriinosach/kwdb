package commands

import (
	"context"
)

const CommandPing = "ping"

type PingCommand struct {
	name string
}

func NewPingCommand() *PingCommand {
	return &PingCommand{
		name: CommandPing,
	}
}

func (c *PingCommand) IsWritable(ctx context.Context) bool {
	return false
}

func (c *PingCommand) SetArgs(ctx context.Context, args *arguments) {
	return
}

func (c *PingCommand) CheckArgs(ctx context.Context, args *arguments) bool {
	return true
}

func (c *PingCommand) Name() string {

	return c.name
}

func (c *PingCommand) Execute(ctx context.Context) (string, error) {
	return "pong", nil
}
