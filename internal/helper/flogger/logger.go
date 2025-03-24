package flogger

import (
	"fmt"
	"kwdb/app"
	"os"
	"time"
)

type FileLogger struct {
}

func (f FileLogger) Write(mes []byte) (n int, err error) {
	y, m, d := time.Now().Date()
	logFileDate := fmt.Sprintf("log-%d-%d-%d", d, m, y)

	message := string(mes[:])

	filePath := app.Config.LogPath + "/" + logFileDate + ".txt"

	var file *os.File

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Create the file
		file, err = os.Create(filePath)
		if err != nil {
			// TODO ошибка логирования
			panic("ошибка создания файла логирования:" + err.Error())
			return 0, nil
		}
	} else {
		// Open the file in append mode
		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			// TODO ошибка логирования
			fmt.Println("ошибка открытия файла логирования:", err)
			return 0, err
		}

	}
	defer file.Close()

	_, err = file.WriteString(message)
	if err != nil {
		// TODO ошибка логирования
		fmt.Println("ошибка записи в файл логирования:", err)
		return 0, err
	}

	return 0, nil
}

func Write(mes string, stream string) {

	if stream == "" {
		stream = "log"
	}

	FileLogger{}.Write([]byte(mes))
}
