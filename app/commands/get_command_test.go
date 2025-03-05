package commands_test

import (
	"context"
	"kwdb/app/commands"
	"testing"
)

func TestGetCommand(t *testing.T) {

	args := new(commands.CommandArguments)

	args.CmdName = "SET"
	args.Key = "1"
	args.Value = "1"
	args.TTL = 100

	ctx := context.Background()
	// TODO: как выбирать инкапсулированные методы / selectCommand
	cmd := commands.NewGetCommand()
	
	result := cmd.CheckArgs(ctx, args)

	if !result {
		t.Errorf("Команда %s отвергла необходимые аргументы: %v.", args.CmdName, args)
	}
}
