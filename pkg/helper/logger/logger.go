package logger

import (
	"fmt"
	"kwdb/app"
	"os"
	"time"
)

var logFileName_base = "log"

func Write(message string) {

	y, m, d := time.Now().Date()
	logFileDate := fmt.Sprintf("-%d-%d-%d", d, m, y)

	filePath := app.Config.LogPath + "/" + logFileName_base + logFileDate + ".txt"

	var file *os.File
	var err error

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Create the file
		file, err = os.Create(filePath)
		if err != nil {
			// TODO ошибка логирования
			panic("ошибка создания файла логирования:" + err.Error())
			return
		}
	} else {
		// Open the file in append mode
		file, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			// TODO ошибка логирования
			fmt.Println("ошибка открытия файла логирования:", err)
			return
		}

	}
	defer file.Close()

	_, err = file.WriteString(message + "\n")
	if err != nil {
		// TODO ошибка логирования
		fmt.Println("ошибка записи в файл логирования:", err)
		return
	}

}
