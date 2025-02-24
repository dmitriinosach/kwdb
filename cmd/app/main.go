package main

import (
	"fmt"
	"kwdb/app"
	"kwdb/app/api/http"
	"kwdb/app/api/tcp"
	"kwdb/app/storage"
)

func main() {

	//загрузка настроек
	_, err := app.InitConfigs()

	if err != nil {
		panic("не установленны базовые настройки: " + err.Error())
	}

	err = storage.Init(app.Config.DRIVER)
	if err != nil {
		fmt.Println("Ошибка инициализации хранилища: ", err)
		return
	}

	//задача чистильщика
	//go workers.CleanerRun()

	go http.Serve()

	tcp.Serve()
}
