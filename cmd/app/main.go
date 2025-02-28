package main

import (
	"kwdb/app"
	"kwdb/app/api/http"
	"kwdb/app/api/tcp"
	"kwdb/app/helpers"
	"kwdb/app/storage"
	"kwdb/app/workers/cleaner"
	"os"
)

func main() {

	//Консольный информатор
	go helpers.ConsoleInformer()

	//загрузка настроек
	_, err := app.InitConfigs()

	if err != nil {
		helpers.InfChan <- "Ошибка чтения настроек:" + err.Error()
		os.Exit(-1)
	}

	helpers.InfChan <- "Настройки приложения загружены"

	err = storage.Init(app.Config.DRIVER, app.Config.PARTITIONS)

	if err != nil {
		helpers.InfChan <- "Ошибка инициализации хранилища:" + err.Error()
		os.Exit(-1)
	}

	helpers.InfChan <- "Хранилище инициализировано:"
	helpers.InfChan <- "Состояние хранилища: \n" + storage.Storage.Info()

	//задача чистильщика
	go cleaner.Run(app.Config.PARTITIONS)

	go http.Serve()

	tcp.Serve()
}
