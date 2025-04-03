package main

import (
	"context"
	"kwdb/app"
	"kwdb/app/api/http"
	"kwdb/app/api/tcp"
	"kwdb/app/storage"
	cprntr "kwdb/internal/helper/color_printer"
	"os"
)

var Channels map[string]chan string

// загрузка настроек
func loadConfigs() {
	err := app.InitConfigs()

	if err != nil {
		panic("Ошибка загрузки настроек:" + err.Error())
	}

	app.InfChan <- cprntr.Green + "Настройки загружены" + cprntr.Reset
}

// Создание хранилища
func runStorage(ctx context.Context) {
	err := storage.Init(app.Config.Driver, app.Config.Partitions)

	if err != nil {
		app.InfChan <- "Ошибка инициализации хранилища:" + err.Error()
		os.Exit(-1)
	}

	app.InfChan <- cprntr.Green + "Хранилище инициализировано\n" + cprntr.Reset
}

func runListeners(ctx context.Context) {
	go http.Serve(ctx)

	go tcp.Serve(ctx)
}
