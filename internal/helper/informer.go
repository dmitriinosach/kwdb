package helper

import (
	"fmt"
	"kwdb/internal/helper/flogger"
	"log/slog"
	"time"
)

var InfChan = make(chan string, 1)

func ConsoleInformer() {
	for message := range InfChan {
		handler := slog.NewJSONHandler(flogger.FileLogger{}, &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelInfo,
		})

		logger := slog.New(handler)
		logger.Info(message)

		message = "[" + time.Now().Format("2006-01-02 15:04:05") + "] " + message

		fmt.Println(message)
	}
}
