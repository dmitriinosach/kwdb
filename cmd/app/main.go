package main

import (
	"kwdb/app"
	"kwdb/app/api/http"
	"kwdb/app/api/tcp"
	"kwdb/app/storage"
	"kwdb/app/workers/cleaner"
	"kwdb/pkg/helper"
	"os"
)

func main() {

	//Консольный информатор
	go helper.ConsoleInformer()

	//загрузка настроек
	_, err := app.InitConfigs()

	if err != nil {
		helper.InfChan <- "Ошибка чтения настроек:" + err.Error()
		os.Exit(-1)
	}

	helper.InfChan <- "Настройки приложения загружены"

	err = storage.Init(app.Config.DRIVER, app.Config.PARTITIONS)

	if err != nil {
		helper.InfChan <- "Ошибка инициализации хранилища:" + err.Error()
		os.Exit(-1)
	}

	helper.InfChan <- "Хранилище инициализировано:"
	helper.InfChan <- "Состояние хранилища: \n" + storage.Storage.Info()

	//задача чистильщика
	go cleaner.Run(app.Config.PARTITIONS)

	go http.Serve()

	tcp.Serve()
}
