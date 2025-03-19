package main

import (
	"kwdb/app"
	"kwdb/app/workers/cleaner"
	"kwdb/pkg/helper"
)

func main() {

	//загрузка настроек
	loadConfigs()
	//Консольный информатор
	go helper.ConsoleInformer()

	helper.InfChan <- "Настройки загружены"

	helper.InfChan <- "Запуск..."

	//Создание хранилища
	runStorage()

	//задача чистильщика
	go cleaner.Run(app.Config.Partitions)

	//Запуск слушателей
	runListeners()
}
