package commands

import (
	"context"
	"kwdb/app/errorpkg"
	"kwdb/app/workers"
)

var List = map[string]CommandInterface{
	// Команды работы с базой данных
	CommandGet:     NewGetCommand(),
	CommandSet:     NewSetCommand(),
	CommandDelete:  NewDeleteCommand(),
	CommandInfo:    NewInfoCommand(),
	CommandRestore: NewRestoreCommand(),
	CommandLookUp:  NewLookUpCommand(),

	//Команды управления и дебага
	CommandStatus: NewStatusCommand(),
	CommandPing:   NewPingCommand(),
}

type CommandInterface interface {
	Name() string
	Execute(ctx context.Context) (string, error)
	CheckArgs(ctx context.Context, args *arguments) bool
	SetArgs(ctx context.Context, args *arguments)
	IsWritable(ctx context.Context) bool
}

func setupCommand(ctx context.Context, message string) (CommandInterface, error) {

	args, err := newArgsFromString(message)

	if err != nil {
		return nil, errorpkg.ErrCmdLineParser
	}

	cmd := selectCommand(args)

	if cmd == nil {
		return nil, errorpkg.ErrCmdNotFound
	}

	if !cmd.CheckArgs(ctx, args) {
		return nil, errorpkg.ErrCmdArguments
	}

	cmd.SetArgs(ctx, args)

	return cmd, nil
}

func SetAndRun(ctx context.Context, message string) (string, error) {
	cmd, err := setupCommand(ctx, message)

	if err != nil {
		return "", err
	}

	execute, err := cmd.Execute(ctx)
	if err != nil {
		return "", err
	}

	if cmd.IsWritable(ctx) {
		go workers.Write(message)
	}

	return execute, nil
}

func selectCommand(args *arguments) CommandInterface {
	if List[args.CmdName] == nil {
		return nil
	}

	return List[args.CmdName]
}
