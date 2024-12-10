package commands

import (
	"fmt"
)

func List() []CommandInterface {
	commandsList := []CommandInterface{}

	commandsList = append(
		commandsList,
		&GetCommand{"GET", CommandArguments{}},
		&SetCommand{"SET", CommandArguments{}},
		&InfoCommand{"INFO", CommandArguments{}},
		&RestoreCommand{"RESTORE", CommandArguments{}})

	return commandsList
}

func SetupCommand(message string) (CommandInterface, error) {
	args := Parce(message)

	cmd := selectCommand(args)

	if cmd == nil {

		return nil, fmt.Errorf("команда не найдена")
	}

	if !cmd.CheckArgs(args) {
		return nil, fmt.Errorf("отсутствуют необходимые аргуметы")
	}

	cmd.SetArgs(args)

	return cmd, nil
}

func selectCommand(args CommandArguments) CommandInterface {

	var command CommandInterface
	for _, cmd := range List() {
		if cmd.Name() == args.Name {
			command = cmd
			break
		}
	}

	if command == nil {
		return nil
	}
	//fmt.Errorf("команда не найдена")
	return command
}
