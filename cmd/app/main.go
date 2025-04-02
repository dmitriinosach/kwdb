package main

import (
	"context"
	"fmt"
	"kwdb/internal/helper/flogger"
	"kwdb/internal/helper/informer"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(1)
	ctx, shutDown := context.WithCancel(context.Background())
	loadConfigs()

	flogger.Init()

	go informer.Informer(ctx)

	//загрузка настроек

	//Создание хранилища
	runStorage(ctx)

	//Запуск слушателей
	runListeners(ctx)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Создаем горутину для обработки сигналов
	go func() {

		sig := <-sigs

		if sig == syscall.SIGTERM || sig == syscall.SIGINT {
			shutDown()
			fmt.Println("Shutting Down")
		}
		wg.Done()
	}()

	wg.Wait()
}
