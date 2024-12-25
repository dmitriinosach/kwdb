package app

import (
	"context"
	"github.com/pkg/errors"
	"kwdb/app/commands"
	"kwdb/app/storage"
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

	storage.Storage.Lock()
	result, err := cmd.Execute(ctx)
	storage.Storage.Unlock()

	return result, err
}
