package commands

import (
	"fmt"
	"kwdb/app/storage"
)

type LookUpCommand struct {
	name       string
	Args       CommandArguments
	isWritable bool
}

func (command *LookUpCommand) IsWritable() bool {
	return command.isWritable
}

func (command *LookUpCommand) CheckArgs(args CommandArguments) bool {
	return true
}

func (command *LookUpCommand) Execute() (string, error) {

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

func (command *LookUpCommand) SetArgs(args CommandArguments) {
	command.Args = args
}
