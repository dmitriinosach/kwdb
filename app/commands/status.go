package commands

import (
	"context"
	"fmt"
	"kwdb/app/storage"
	"kwdb/internal/helper"
	"runtime"
	"strconv"
	"time"
)

const CommandStatus = "status"

type StatusCommand struct {
	name       string
	Args       *arguments
	isWritable bool
}

func NewStatusCommand() *StatusCommand {
	return &StatusCommand{
		name:       CommandStatus,
		Args:       new(arguments),
		isWritable: false,
	}
}

func (c *StatusCommand) IsWritable(ctx context.Context) bool {
	return c.isWritable
}

func (c *StatusCommand) SetArgs(ctx context.Context, args *arguments) {
	c.Args = args
}

func (c *StatusCommand) CheckArgs(ctx context.Context, args *arguments) bool {
	return true
}

func (c *StatusCommand) Name() string {

	return c.name
}

func (c *StatusCommand) Execute(ctx context.Context) (string, error) {

	// ваш код
	duration := storage.Status.Uptime()
	minutes := (duration % time.Hour) / time.Minute
	seconds := (duration % time.Minute) / time.Second

	status := "coroutines:" + strconv.Itoa(runtime.NumGoroutine()) + "\n"
	status += "cores:" + strconv.Itoa(runtime.NumCPU()) + " \n"
	status += "driver:" + storage.Status.DriverName + " \n"
	status += "hitrate:" + storage.Status.HitRate() + " \n"
	status += "lifetime: " + fmt.Sprintf("%dч. %dм. %dс.\n", int(duration.Hours()), minutes, seconds)
	status += "" + helper.MemStatInfo() + "\n"

	return status, nil
}
