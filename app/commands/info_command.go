package commands

import (
	"context"
	"strconv"
	"time"

	"kwdb/app/storage"
)

const CommandInfo = "INFO"

type InfoCommand struct {
	name       string
	Args       *CommandArguments
	isWritable bool
}

func NewInfoCommand() *InfoCommand {
	return &InfoCommand{
		name:       CommandInfo,
		Args:       new(CommandArguments),
		isWritable: false,
	}
}

func (command *InfoCommand) IsWritable(ctx context.Context) bool {
	return command.isWritable
}

func (command *InfoCommand) SetArgs(ctx context.Context, args *CommandArguments) {
	command.Args = args
}

func (command *InfoCommand) CheckArgs(ctx context.Context, args *CommandArguments) bool {
	return true
}

func (command *InfoCommand) Name() string {

	return command.name
}

func (command *InfoCommand) Execute(ctx context.Context) (string, error) {

	info := "driver:" + storage.Storage.GetDriver() + "\n"
	info += "Length:" + strconv.Itoa(len(storage.Storage.GetVaultMap())) + "\n"
	info += "Uptime:" + time.Unix(time.Now().Unix()-storage.Started.Unix(), 0).Format("05 04 15") + "\n"

	return info, nil
}
