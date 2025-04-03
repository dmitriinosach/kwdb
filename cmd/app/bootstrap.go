package main

import (
	"context"
	"kwdb/app"
	"kwdb/app/api/http"
	"kwdb/app/api/tcp"
	"kwdb/app/storage"
	cprntr "kwdb/internal/helper/color_printer"
	"kwdb/internal/helper/informer"
	"os"
)

// загрузка настроек
func loadConfigs() {
	err := app.InitConfigs()

	if err != nil {
		panic("Ошибка загрузки настроек:" + err.Error())
	}
	informer.InfChan <- cprntr.Green + "Настройки загружены" + cprntr.Reset
}

// Создание хранилища
func runStorage(ctx context.Context) {
	err := storage.Init(app.Config.Driver, app.Config.Partitions)

	if err != nil {
		informer.InfChan <- "Ошибка инициализации хранилища:" + err.Error()
		os.Exit(-1)
	}

	informer.InfChan <- cprntr.Green + "Хранилище инициализировано\n" + cprntr.Reset
}

func runListeners(ctx context.Context) {
	go http.Serve(ctx)

	go tcp.Serve(ctx)
}
