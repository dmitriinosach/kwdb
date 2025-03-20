package workers

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// TODO: переделать, долен решать другую задачу
func Write(text string) {

	file, err := os.OpenFile("./data/backup/wal1.txt", os.O_APPEND, 066)

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.WriteString(text + "\n")
	if err != nil {
		return
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
