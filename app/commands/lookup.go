package commands

import (
	"strings"
)

const CommandLookUp = "LOOKUP"

type LookUpCommand struct {
	name       string
	Args       *arguments
	isWritable bool
}

func NewLookUpCommand() CommandInterface {
	return &LookUpCommand{
		name:       CommandLookUp,
		Args:       new(arguments),
		isWritable: false,
	}
}

func (command *LookUpCommand) IsWritable() bool {
	return command.isWritable
}

func (command *LookUpCommand) CheckArgs() bool {
	return true
}

func (command *LookUpCommand) Execute() ([]byte, error) {

	reply := strings.Builder{}

	/*	for k, v := range storage.Storage.GetVaultMap() {
		if len(v.Value) > 10 {
			reply.WriteString(k + ":")
			reply.Write(v.Value[:7])
			reply.WriteString("...")
		} else {
			reply.WriteString(k + ":")
			reply.Write(v.Value)
		}
	}*/

	return []byte(reply.String()), nil
}

func (command *LookUpCommand) Name() string {
	return command.name
}

func (command *LookUpCommand) SetArgs(args *arguments) {
	command.Args = args
}
