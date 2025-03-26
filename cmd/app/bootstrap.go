package main

import (
	"kwdb/app"
	"kwdb/app/api/http"
	"kwdb/app/api/tcp"
	"kwdb/app/storage"
	"kwdb/internal/helper/informer"
	"os"
)

// загрузка настроек
func loadConfigs() {

	err := app.InitConfigs()

	if err != nil {
		panic("Ошибка загрузки настроек:" + err.Error())
	}
}

// Создание хранилища
func runStorage() {
	informer.InfChan <- "Создание хранилища..."

	err := storage.Init(app.Config.Driver, app.Config.Partitions)

	if err != nil {
		informer.InfChan <- "Ошибка инициализации хранилища:" + err.Error()
		os.Exit(-1)
	}

	informer.InfChan <- "Хранилище инициализировано:\n" + storage.Storage.Info()
}

func runListeners() {
	go http.Serve()

	tcp.Serve()
}
