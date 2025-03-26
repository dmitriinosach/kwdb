package workers

import (
	"bufio"
	"kwdb/internal/helper/file_system"
	"kwdb/internal/helper/informer"
	"log"
	"os"
)

const backupPath = "./data/backup/wal1.txt"

var backupFile *os.File

// TODO: переделать, долен решать другую задачу
func Write(text string) {

	if backupFile == nil {
		backupFile, _ = file_system.ReadOrCreate(backupPath)
	}

	_, err := backupFile.WriteString(text + "\n")

	if err != nil {
		informer.InfChan <- "Ошибка записи wal"
	}
}

func Backup(commandChan chan string) *bufio.Scanner {
	file, err := os.OpenFile("./data/backup/wal1.txt", os.O_APPEND, 066)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		cmdString := scanner.Text()

		if cmdString != "" {
			commandChan <- cmdString
		}

	}
	close(commandChan)

	return scanner
}
