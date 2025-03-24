package commands

import (
	"context"
	"fmt"
	"kwdb/app/storage"
)

const CommandGet = "GET"

var (
	GetCommandKeyNotFound = fmt.Errorf("ключ не установлен")
)

type GetCommand struct {
	name       string
	Args       *arguments
	isWritable bool
}

func NewGetCommand() *GetCommand {
	return &GetCommand{
		name:       CommandGet,
		Args:       new(arguments),
		isWritable: false,
	}
}

func (c *GetCommand) IsWritable(ctx context.Context) bool {
	return c.isWritable
}

func (c *GetCommand) SetArgs(ctx context.Context, args *arguments) {
	c.Args = args
}

func (c *GetCommand) CheckArgs(ctx context.Context, args *arguments) bool {
	if args.Key == "" {
		return false
	}

	return true
}

func (c *GetCommand) Name() string {

	return c.name
}

func (c *GetCommand) echo() string {

	return c.name
}

func (c *GetCommand) Execute(ctx context.Context) (string, error) {

	value, err := storage.Storage.Get(ctx, c.Args.Key)

	if value == nil {
		return "", err
	}

	return value.Value, nil
}
