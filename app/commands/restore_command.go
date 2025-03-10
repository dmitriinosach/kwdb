package commands

import (
	"context"
	"fmt"
	"kwdb/app/workers"
)

const CommandRestore = "RESTORE"

type RestoreCommand struct {
	name       string
	Args       *Arguments
	isWritable bool
}

func NewRestoreCommand() *RestoreCommand {
	return &RestoreCommand{
		name:       CommandRestore,
		Args:       new(Arguments),
		isWritable: false,
	}
}
func (c *RestoreCommand) CheckArgs(ctx context.Context, args *Arguments) bool {
	return true
}

func (c *RestoreCommand) Execute(ctx context.Context) (string, error) {

	ch := make(chan string)

	go workers.Backup(ch)
	for {
		commandString, ok := <-ch
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

func (c *RestoreCommand) Name() string {
	return c.name
}

func (c *RestoreCommand) SetArgs(ctx context.Context, args *Arguments) {
	c.Args = args
}

func (c *RestoreCommand) IsWritable(ctx context.Context) bool {
	return c.isWritable
}
