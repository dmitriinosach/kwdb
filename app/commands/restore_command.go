package commands

import (
	"context"
	"fmt"

	"kwdb/app/workers"
)

const CommandRestore = "RESTORE"

type RestoreCommand struct {
	name       string
	Args       *CommandArguments
	isWritable bool
}

func NewRestoreCommand() *RestoreCommand {
	return &RestoreCommand{
		name:       CommandRestore,
		Args:       new(CommandArguments),
		isWritable: false,
	}
}
func (command *RestoreCommand) CheckArgs(ctx context.Context, args *CommandArguments) bool {
	return true
}

func (command *RestoreCommand) Execute(ctx context.Context) (string, error) {

	c := make(chan string)

	go workers.Backup(c)
	for {
		commandString, ok := <-c
		if ok == false {
			if commandString == "" {
				fmt.Println("Done")
			} else {
				fmt.Println(commandString, ok, "<-- loop broke!")
			}
			break // exit break loop

		} else {
			cmd, _ := SetupCommand(ctx, commandString)
			_, err := cmd.Execute(ctx)
			if err != nil {
				return "", err
			}
		}
	}

	return "", nil
}

func (command *RestoreCommand) Name() string {
	return command.name
}

func (command *RestoreCommand) SetArgs(ctx context.Context, args *CommandArguments) {
	command.Args = args
}

func (command *RestoreCommand) IsWritable(ctx context.Context) bool {
	return command.isWritable
}
