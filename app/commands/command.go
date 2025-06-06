package commands

import (
	"kwdb/app/backup"
	"kwdb/app/errorpkg"
	"kwdb/app/storage"
)

var List = map[string]func() CommandInterface{
	// Команды работы с базой данных
	CommandGet:     NewGetCommand,
	CommandSet:     NewSetCommand,
	CommandDelete:  NewDeleteCommand,
	CommandInfo:    NewInfoCommand,
	CommandRestore: NewRestoreCommand,
	CommandLookUp:  NewLookUpCommand,

	//Команды управления и дебага
	CommandStatus:   NewStatusCommand,
	CommandPing:     NewPingCommand,
	CommandFlush:    NewFlushCommand,
	CommandTruncate: NewTruncateCommand,
}

type CommandInterface interface {
	Name() string
	Execute() ([]byte, error)
	CheckArgs() bool
	SetArgs(args *arguments)
	IsWritable() bool
}

func setupCommand(message []byte) (CommandInterface, error) {

	args, err := newArgsFromString(string(message))

	if err != nil {
		return nil, errorpkg.ErrCmdLineParser
	}

	var cmd CommandInterface

	if List[args.CmdName] != nil {
		cmd = List[args.CmdName]()
	} else {
		return nil, errorpkg.ErrCmdNotFound
	}

	cmd.SetArgs(args)

	if !cmd.CheckArgs() {
		return nil, errorpkg.ErrCmdArguments
	}

	return cmd, nil
}

func SetAndRun(message []byte) ([]byte, error) {
	cmd, err := setupCommand(message)

	if err != nil {
		return []byte{}, err
	}

	execute, err := cmd.Execute()
	if err != nil {
		return []byte{}, err
	}

	if cmd.IsWritable() && !storage.Status.Restoring.Load() {
		go backup.Write(message)
	}

	return execute, nil
}
