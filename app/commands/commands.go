package commands

import (
	"fmt"
)

var List = []CommandInterface{
	&GetCommand{"GET", CommandArguments{}, false},
	&SetCommand{"SET", CommandArguments{}, true},
	&InfoCommand{"INFO", CommandArguments{}, false},
	&RestoreCommand{"RESTORE", CommandArguments{}, false},
	&DeleteCommand{"DELETE", CommandArguments{}, true},
	&UpdateCommand{"UPDATE", CommandArguments{}, true},
	&LookUpCommand{"LOOKUP", CommandArguments{}, false},
}

type CommandInterface interface {
	Name() string
	Execute() (string, error)
	CheckArgs(args CommandArguments) bool
	SetArgs(args CommandArguments)
	IsWritable() bool
}

type CommandArguments struct {
	Name  string
	Key   string
	Value string
	TTL   int
}

func SetupCommand(message string) (CommandInterface, error) {

	args, ok := Parse(message)

	if ok != nil {
		return nil, fmt.Errorf("ошибка разрабора запроса")
	}

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
	for _, cmd := range List {
		if cmd.Name() == args.Name {
			command = cmd
			break
		}
	}

	if command == nil {
		return nil
	}

	return command
}
