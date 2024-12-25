package commands

import (
	"fmt"
	"kwdb/app/workers"
)

type RestoreCommand struct {
	name       string
	Args       CommandArguments
	isWritable bool
}

func (command *RestoreCommand) CheckArgs(args CommandArguments) bool {
	return true
}

func (command *RestoreCommand) Execute() (string, error) {

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
			cmd, _ := SetupCommand(commandString)
			_, err := cmd.Execute()
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

func (command *RestoreCommand) SetArgs(args CommandArguments) {
	command.Args = args
}

func (command *RestoreCommand) IsWritable() bool {
	return command.isWritable
}
