package commands

import (
	"context"
	"kwdb/app/storage"
)

const CommandLookUp = "LOOKUP"

type LookUpCommand struct {
	name       string
	Args       *CommandArguments
	isWritable bool
}

func NewLookUpCommand() *LookUpCommand {
	return &LookUpCommand{
		name:       CommandLookUp,
		Args:       new(CommandArguments),
		isWritable: false,
	}
}

func (command *LookUpCommand) IsWritable(ctx context.Context) bool {
	return command.isWritable
}

func (command *LookUpCommand) CheckArgs(ctx context.Context, args *CommandArguments) bool {
	return true
}

func (command *LookUpCommand) Execute(ctx context.Context) (string, error) {

	reply := ""

	for k, v := range storage.Storage.GetVaultMap() {
		if len(v.Value) > 10 {
			reply += k + ":" + v.Value[:7] + "..."
		} else {
			reply += k + ":" + v.Value
		}
	}

	return reply, nil
}

func (command *LookUpCommand) Name() string {
	return command.name
}

func (command *LookUpCommand) SetArgs(ctx context.Context, args *CommandArguments) {
	command.Args = args
}
