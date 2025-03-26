package main

import (
	"kwdb/app"
	"kwdb/app/workers/cleaner"
	"kwdb/internal/helper/flogger"
	"kwdb/internal/helper/informer"
)

func main() {

	//загрузка настроек
	loadConfigs()

	flogger.Init()
	
	//Консольный информатор
	go informer.Informer()

	informer.InfChan <- "Настройки загружены"

	informer.InfChan <- "Запуск..."

	//Создание хранилища
	runStorage()

	//задача чистильщика
	go cleaner.Run(app.Config.Partitions)

	//Запуск слушателей
	runListeners()
}
