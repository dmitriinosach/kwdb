package commands

import (
	"context"
	"fmt"

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

	for k, v := range storage.Storage.GetVaultMap() {
		if len(v.Value) > 10 {
			fmt.Println(k, ":", v.Value[:7], "...")
		} else {
			fmt.Println(k, ":", v.Value)
		}
	}

	return "", nil
}

func (command *LookUpCommand) Name() string {
	return command.name
}

func (command *LookUpCommand) SetArgs(ctx context.Context, args *CommandArguments) {
	command.Args = args
}
