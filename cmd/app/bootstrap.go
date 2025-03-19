package main

import (
	"kwdb/app"
	"kwdb/app/api/http"
	"kwdb/app/api/tcp"
	"kwdb/app/storage"
	"kwdb/pkg/helper"
	"os"
)

// загрузка настроек
func loadConfigs() {

	_, err := app.InitConfigs()

	if err != nil {
		panic("Ошибка загрузки настроек:" + err.Error())
	}
}

// Создание хранилища
func runStorage() {
	helper.InfChan <- "Создание хранилища..."

	err := storage.Init(app.Config.Driver, app.Config.Partitions)

	if err != nil {
		helper.InfChan <- "Ошибка инициализации хранилища:" + err.Error()
		os.Exit(-1)
	}

	helper.InfChan <- "Хранилище инициализировано:\n" + storage.Storage.Info()
}

func runListeners() {
	go http.Serve()

	tcp.Serve()
}
