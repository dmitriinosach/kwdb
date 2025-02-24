package logger

import (
	"fmt"
	"os"
)

func Write(message string) {
	file, err := os.OpenFile("./logs/log.txt", os.O_APPEND, 066)

	if err != nil {
		fmt.Println("ошибка открытия файла:", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("ошибка закрытия файла")
		}
	}(file)

	_, err = file.WriteString(message + "\n")
	if err != nil {
		return
	}
}
