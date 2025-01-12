package commands

import (
	"context"
	"github.com/pkg/errors"
	"kwdb/app/storage"
)

const CommandGet = "GET"

var (
	GetCommandNotFound = errors.New("ключ не установлен")
)

type GetCommand struct {
	name       string
	Args       *CommandArguments
	isWritable bool
}

func NewGetCommand() *GetCommand {
	return &GetCommand{
		name:       CommandGet,
		Args:       new(CommandArguments),
		isWritable: false,
	}
}

func (command *GetCommand) IsWritable(ctx context.Context) bool {
	return command.isWritable
}

func (command *GetCommand) SetArgs(ctx context.Context, args *CommandArguments) {
	command.Args = args
}

func (command *GetCommand) CheckArgs(ctx context.Context, args *CommandArguments) bool {
	if args.Key == "" {
		return false
	}

	return true
}

func (command *GetCommand) Name() string {

	return command.name
}

func (command *GetCommand) Execute(ctx context.Context) (string, error) {

	if !storage.Storage.HasKey(command.Args.Key) {
		return "", GetCommandNotFound
	}

	value, _ := storage.Storage.GetValue(command.Args.Key)

	return value.Value, nil
}
