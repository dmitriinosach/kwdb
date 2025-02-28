package main

import (
	"kwdb/app"
	"kwdb/app/api/http"
	"kwdb/app/api/tcp"
	"kwdb/app/storage"
	"kwdb/app/workers/cleaner"
	"os"
)

func main() {

	//Консольный информатор
	go app.ConsoleInformer()

	//загрузка настроек
	_, err := app.InitConfigs()

	if err != nil {
		app.InfChan <- "Ошибка чтения настроек:" + err.Error()
		os.Exit(-1)
	}

	app.InfChan <- "Настройки приложения загружены"

	err = storage.Init(app.Config.DRIVER, app.Config.PARTITIONS)

	if err != nil {
		app.InfChan <- "Ошибка инициализации хранилища:" + err.Error()
		os.Exit(-1)
	}

	app.InfChan <- "Хранилище инициализировано:"
	app.InfChan <- "Состояние хранилища: \n" + storage.Storage.Info()

	//задача чистильщика
	go cleaner.Run(app.Config.PARTITIONS)

	go http.Serve()

	tcp.Serve()
}
