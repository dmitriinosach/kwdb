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

func TestSetCommand(t *testing.T) {

	args := new(commands.CommandArguments)

	args.CmdName = "SET"
	args.Key = "1"
	args.Value = "1"
	args.TTL = 100

	ctx := context.Background()
	cmd := commands.NewSetCommand()
	result := cmd.CheckArgs(ctx, args)

	if !result {
		t.Errorf("Команда %s отвергла необходимые аргументы: %v.", args.CmdName, args)
	}
}

func TestDeleteCommand(t *testing.T) {

	args := new(commands.CommandArguments)

	args.CmdName = "DELETE"
	args.Key = "1"

	ctx := context.Background()
	cmd := commands.NewDeleteCommand()
	result := cmd.CheckArgs(ctx, args)

	if !result {
		t.Errorf("Команда %s отвергла необходимые аргументы: %v.", args.CmdName, args)
	}
}

func TestInfoCommand(t *testing.T) {

	args := new(commands.CommandArguments)

	args.CmdName = "INFO"

	ctx := context.Background()
	cmd := commands.NewInfoCommand()
	result := cmd.CheckArgs(ctx, args)

	if !result {
		t.Errorf("Команда %s отвергла необходимые аргументы: %v.", args.CmdName, args)
	}
}
