package commands

import (
	"context"
	"github.com/pkg/errors"
)

var (
	ErrCommandNotFound   = errors.New("команда не найдена")
	ErrCommandLineParser = errors.New("ошибка разбора аргументов")
	ErrCommandArguments  = errors.New("отсутствуют необходимые аргуметы")
)

var List = map[string]CommandInterface{
	CommandGet:     NewGetCommand(),
	CommandSet:     NewSetCommand(),
	CommandDelete:  NewDeleteCommand(),
	CommandInfo:    NewInfoCommand(),
	CommandRestore: NewRestoreCommand(),
	CommandLookUp:  NewLookUpCommand(),
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
		return nil, ErrCommandLineParser
	}

	cmd := selectCommand(args)

	if cmd == nil {
		return nil, ErrCommandNotFound
	}

	if !cmd.CheckArgs(ctx, args) {
		return nil, ErrCommandArguments
	}

	cmd.SetArgs(ctx, args)

	return cmd, nil
}

func selectCommand(args *CommandArguments) CommandInterface {
	if List[args.Name] == nil {
		return nil
	}

	return List[args.Name]
}
