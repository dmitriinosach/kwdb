package file_system

import (
	"os"
)

func ReadOrCreate(filePath string) (*os.File, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic("ошибка открытия файла:" + filePath)
	}
	return file, err
}
