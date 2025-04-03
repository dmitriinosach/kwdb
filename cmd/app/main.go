package main

import (
	"context"
	"kwdb/app"
	"kwdb/internal/helper/flogger"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}

	ctx, shutDown := context.WithCancel(context.Background())

	go app.ChanHandler(ctx, wg, shutDown)

	loadConfigs()

	flogger.Init()

	//Создание хранилища
	runStorage(ctx)

	//Запуск слушателей
	runListeners(ctx)

	wg.Wait()
}
