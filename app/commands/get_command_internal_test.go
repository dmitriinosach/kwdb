package commands

import (
	"context"
	"testing"
)

//пакет не меняем

func TestGetCommand(t *testing.T) {

	args := new(CommandArguments)

	args.CmdName = "SET"
	args.Key = "1"
	args.Value = "1"
	args.TTL = 100

	ctx := context.Background()
	// TODO: как выбирать инкапсулированные методы / selectCommand
	cmd := NewGetCommand()

	//инкапсулированный тест
	cmd.echo()

	result := cmd.CheckArgs(ctx, args)

	if !result {
		t.Errorf("Команда %s отвергла необходимые аргументы: %v.", args.CmdName, args)
	}
}
