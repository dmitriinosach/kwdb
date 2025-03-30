package commands

import (
	"fmt"
	"kwdb/app/storage"
)

const CommandGet = "GET"

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

func (c *GetCommand) IsWritable() bool {
	return c.isWritable
}

func (c *GetCommand) SetArgs(args *arguments) {
	c.Args = args
}

func (c *GetCommand) CheckArgs() bool {
	if c.Args.Key == "" {
		return false
	}

	return true
}

func (c *GetCommand) Name() string {

	return c.name
}

func (c *GetCommand) Execute() (string, error) {

	cell, err := storage.Storage.Get(c.Args.Key)
	fmt.Printf("%v", cell)
	if cell == nil {
		storage.Status.Metrics.Miss()

		return "", err
	}

	storage.Status.Metrics.Hit()

	return cell.Value, nil
}
