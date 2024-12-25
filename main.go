package main

import (
	"fmt"

	"kwdb/app"
	"kwdb/app/api"
	"kwdb/app/storage"
	"kwdb/app/workers"
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

	//Отельная рутина для работы CLI
	go api.HandleCLI()

	//задача чистильщика
	go workers.CleanerRun()

	fmt.Printf("Приложение запущено\n")

	app.ServeTCP()
}
