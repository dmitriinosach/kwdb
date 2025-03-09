package helper

import (
	"fmt"
	"kwdb/pkg/helper/logger"
	"time"
)

var InfChan = make(chan string, 1)

func ConsoleInformer() {
	for message := range InfChan {
		message = "[" + time.Now().Format("2006-01-02 15:04:05") + "] " + message
		fmt.Println(message)
		logger.Write(message)
	}
}
