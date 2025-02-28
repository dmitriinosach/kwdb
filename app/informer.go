package app

import (
	"fmt"
	"kwdb/app/logger"
	"time"
)

var InfChan = make(chan string, 1)

func ConsoleInformer() {

	InfChan <- "Информер запущен"

	for message := range InfChan {
		message = "[" + time.Now().Format("2006-01-02 15:04:05") + "] " + message
		fmt.Println(message)
		logger.Write(message)
	}
}
