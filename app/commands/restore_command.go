package commands

import "kwdb/app/wal"

type RestoreCommand struct {
	name string
	Args CommandArguments
}

func (command *RestoreCommand) CheckArgs(args CommandArguments) bool {
	return true
}

func (command *RestoreCommand) Execute() (string, error) {

	wal.Backup()

	return "", nil
}

func (command *RestoreCommand) Name() string {
	return command.name
}

func (command *RestoreCommand) SetArgs(args CommandArguments) {
	command.Args = args
}
