package main

import (
	"context"
	"fmt"
	"kwdb/app"
	"kwdb/app/api"
	"kwdb/app/commands"
	"kwdb/app/storage"
	"kwdb/app/workers"
	"net/http"
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

	go api.HandleCLI()

	//задача чистильщика
	go workers.CleanerRun()

	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

			ifc := commands.List[commands.CommandInfo]
			ctx := context.Background()
			res, _ := ifc.Execute(ctx)
			fmt.Fprintf(w, res)
		})
		
		http.ListenAndServe("192.168.1.5:83", nil)
	}()

	app.ServeTCP()
}
