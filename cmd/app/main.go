package main

import (
	"fmt"
	"kwdb/app"
	"kwdb/app/api/http"
	"kwdb/app/api/tcp"
	"kwdb/app/storage"
	"kwdb/app/workers/cleaner"
)

func main() {

	//загрузка настроек
	_, err := app.InitConfigs()

	if err != nil {
		panic("не установленны базовые настройки: " + err.Error())
	}

	err = storage.Init(app.Config.DRIVER, app.Config.PARTITIONS)
	if err != nil {
		fmt.Println("Ошибка инициализации хранилища: ", err)
		return
	}

	//задача чистильщика
	go cleaner.Run(app.Config.PARTITIONS)

	go http.Serve()

	tcp.Serve()
}
