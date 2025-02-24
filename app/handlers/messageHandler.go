package api

import (
	"context"
	"github.com/pkg/errors"
	"kwdb/app/commands"
	"kwdb/app/workers"
)

func HandleMessage(ctx context.Context, message string) (string, error) {
	cmd, err := commands.SetupCommand(ctx, message)

	if err != nil {
		return "", errors.Wrap(err, "ошибка установки команды")
	}

	if cmd.IsWritable(ctx) {
		go workers.Write(message)
	}

	result, err := cmd.Execute(ctx)

	return result, err
}
