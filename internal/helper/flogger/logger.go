package flogger

import (
	"fmt"
	"kwdb/app"
	"kwdb/internal/helper/file_system"
	"os"
	"time"
)

type FileLogger struct {
	File *os.File
}

var logger *FileLogger

var file *os.File

func Init() {
	getLogFile()
}

func (f FileLogger) Write(mes []byte) (n int, err error) {
	message := string(mes[:])

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

func getLogFile() {
	y, m, d := time.Now().Date()
	logFileDate := fmt.Sprintf("log-%d-%d-%d", d, m, y)

	filePath := app.Config.LogPath + "/" + logFileDate + ".txt"
	file, _ = file_system.ReadOrCreate(filePath)
}
