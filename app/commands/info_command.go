package commands

import (
	"kwdb/app/storage"
	"strconv"
	"time"
)

type InfoCommand struct {
	name       string
	Args       CommandArguments
	isWritable bool
}

func (command *InfoCommand) IsWritable() bool {
	return command.isWritable
}

func (command *InfoCommand) SetArgs(args CommandArguments) {
	command.Args = args
}

func (command *InfoCommand) CheckArgs(args CommandArguments) bool {
	return true
}

func (command *InfoCommand) Name() string {

	return command.name
}

func (command *InfoCommand) Execute() (string, error) {

	info := "Driver:" + storage.Storage.GetDriver() + "\n"
	info += "Length:" + strconv.Itoa(len(storage.Storage.GetVaultMap())) + "\n"
	info += "Uptime:" + time.Unix(time.Now().Unix()-storage.Started.Unix(), 0).Format("05 04 15") + "\n"

	return info, nil
}

func (command *SetCommand) InfoCommand() bool {
	return command.isWritable
}
