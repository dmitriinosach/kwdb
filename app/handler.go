package app

import (
	"fmt"
	"kwdb/app/commands"
	"kwdb/app/storage"
	"kwdb/app/workers"
)

func HandleMessage(message string) (string, error) {
	cmd, err := commands.SetupCommand(message)

	if err != nil || cmd == nil {
		return "", fmt.Errorf("ошибка установки команды: %v", err)
	}

	if cmd.IsWritable() {
		go workers.Write(message)
	}

	storage.Storage.Lock()
	result, err := cmd.Execute()
	storage.Storage.Unlock()

	return result, err
}
