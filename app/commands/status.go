package commands

import (
	"fmt"
	"kwdb/app/storage"
	"kwdb/internal/helper"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const CommandStatus = "status"

type StatusCommand struct {
	name       string
	Args       *arguments
	isWritable bool
}

func NewStatusCommand() CommandInterface {
	return &StatusCommand{
		name:       CommandStatus,
		Args:       new(arguments),
		isWritable: false,
	}
}

func (c *StatusCommand) IsWritable() bool {
	return c.isWritable
}

func (c *StatusCommand) SetArgs(args *arguments) {
	c.Args = args
}

func (c *StatusCommand) CheckArgs() bool {
	return true
}

func (c *StatusCommand) Name() string {

	return c.name
}

func (c *StatusCommand) Execute() ([]byte, error) {

	duration := storage.Status.Uptime()
	minutes := (duration % time.Hour) / time.Minute
	seconds := (duration % time.Minute) / time.Second

	sb := strings.Builder{}
	sb.WriteString("coroutines:" + strconv.Itoa(runtime.NumGoroutine()) + "\n")
	sb.WriteString("cores:" + strconv.Itoa(runtime.NumCPU()) + " \n")
	sb.WriteString("driver:" + storage.Status.DriverName + " \n")
	sb.WriteString("hitrate:" + storage.Status.HitRate() + " \n")
	sb.WriteString("lifetime: " + fmt.Sprintf("%dч. %dм. %dс.\n", int(duration.Hours()), minutes, seconds))
	sb.WriteString("" + helper.MemStatInfo() + "\n")

	status := sb.String()

	return []byte(status), nil
}
