package commands

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
)

var (
	ErrCommandNotFound = errors.New("команда не найдена")
)

var list = []CommandInterface{
	NewGetCommand(),
	NewSetCommand(),
	NewDeleteCommand(),
	NewInfoCommand(),
	NewRestoreCommand(),
	NewLookUpCommand(),
}

type CommandInterface interface {
	Name() string
	Execute(ctx context.Context) (string, error)
	CheckArgs(ctx context.Context, args *CommandArguments) bool
	SetArgs(ctx context.Context, args *CommandArguments)
	IsWritable(ctx context.Context) bool
}

type CommandArguments struct {
	Name  string
	Key   string
	Value string
	TTL   int
}

func SetupCommand(ctx context.Context, message string) (CommandInterface, error) {
	args, err := Parse(message)
	if err != nil {
		return nil, errors.Wrap(err, "ошибка парсинга аргументов")
	}

	cmd := selectCommand(args)

	if !cmd.CheckArgs(ctx, args) {
		return nil, fmt.Errorf("отсутствуют необходимые аргуметы")
	}

	cmd.SetArgs(ctx, args)

	return cmd, nil
}

func selectCommand(args *CommandArguments) CommandInterface {

	var command CommandInterface
	for _, cmd := range list {
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
