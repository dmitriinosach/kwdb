package commands

import (
	"context"
	"kwdb/pkg/helper"
	"runtime"
	"strconv"
)

const CommandStatus = "status"

type StatusCommand struct {
	name       string
	Args       *CommandArguments
	isWritable bool
}

func NewStatusCommand() *StatusCommand {
	return &StatusCommand{
		name:       CommandStatus,
		Args:       new(CommandArguments),
		isWritable: false,
	}
}

func (c *StatusCommand) IsWritable(ctx context.Context) bool {
	return c.isWritable
}

func (c *StatusCommand) SetArgs(ctx context.Context, args *CommandArguments) {
	c.Args = args
}

func (c *StatusCommand) CheckArgs(ctx context.Context, args *CommandArguments) bool {
	return true
}

func (c *StatusCommand) Name() string {

	return c.name
}

func (c *StatusCommand) Execute(ctx context.Context) (string, error) {
	status := ""
	status += "coroutines:" + strconv.Itoa(runtime.NumGoroutine()) + "\n"
	status += "cores:" + strconv.Itoa(runtime.NumCPU()) + "\n"
	status += "" + helper.MemStatInfo() + "\n"
	return status, nil
}
