package flogger

import (
	"fmt"
	"kwdb/app"
	"kwdb/internal/helper/file_system"
	"log/slog"
	"os"
	"time"
)

type FileLogger struct {
	File *os.File
}

var Flogger *FileLogger

func Init() {

	f, err := getLogFile()

	if err != nil {
		panic("Ошибка создания файла логирования:" + err.Error())
	}

	Flogger = &FileLogger{File: f}

	fmt.Println("Файл логирования инициализирован")
}

func (f *FileLogger) Write(mes []byte) (n int, err error) {
	message := string(mes[:])

	_, err = Flogger.File.WriteString(message)
	if err != nil {
		// TODO ошибка логирования
		fmt.Println("ошибка записи в файл логирования:", err)
		return 0, err
	}

	return 0, nil
}

func (f *FileLogger) WriteString(mes string) {
	f.Write([]byte(mes))
}

func (f *FileLogger) Log(m string) {
	handler := slog.NewJSONHandler(Flogger.File, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	})

	logger := slog.New(handler)
	logger.Info(m)
}

func getLogFile() (file *os.File, err error) {
	y, m, d := time.Now().Date()
	logFileDate := fmt.Sprintf("log-%d-%d-%d", d, m, y)

	fmt.Printf(app.Config.LogPath)
	filePath := app.Config.LogPath + "/" + logFileDate + ".txt"

	f, err := file_system.ReadOrCreate(filePath)

	if err != nil {
		return nil, err
	}
	return f, err
}
