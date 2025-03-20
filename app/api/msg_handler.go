package api

import (
	"context"
	"github.com/pkg/errors"
	"kwdb/app/commands"
	"kwdb/app/workers"
)

func ExecMsg(ctx context.Context, message string) (string, error) {
	cmd, err := commands.SetupCommand(ctx, message)

	if err != nil {
		return "", errors.Wrap(err, "ошибка установки команды")
	}

	// TODO: перенести в команды, их ответственность ?
	if cmd.IsWritable(ctx) {
		go workers.Write(message)
	}

	// TODO: перенести в команды, их ответственность ?
	result, err := cmd.Execute(ctx)

	return result, err
}
