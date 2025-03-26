package file_system

import (
	"os"
)

func ReadOrCreate(filePath string) (*os.File, error) {
	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// Create the file
		file, err := os.Create(filePath)
		if err != nil {
			panic("ошибка создания необходимого файла:" + filePath + " , " + err.Error())
		}
		return file, nil
	} else {
		// Open the file in append mode
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic("ошибка открытия файла:" + filePath)
		}

		return file, err
	}
}
