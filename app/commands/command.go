package commands

import (
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
	Execute() (string, error)
	CheckArgs() bool
	SetArgs(args *arguments)
	IsWritable() bool
}

func setupCommand(message string) (CommandInterface, error) {

	args, err := newArgsFromString(message)

	if err != nil {
		return nil, errorpkg.ErrCmdLineParser
	}

	cmd := selectCommand(args)

	if cmd == nil {
		return nil, errorpkg.ErrCmdNotFound
	}

	if !cmd.CheckArgs() {
		return nil, errorpkg.ErrCmdArguments
	}

	cmd.SetArgs(args)

	return cmd, nil
}

func SetAndRun(message string) (string, error) {
	cmd, err := setupCommand(message)

	if err != nil {
		return "", err
	}

	execute, err := cmd.Execute()
	if err != nil {
		return "", err
	}

	if cmd.IsWritable() {
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
