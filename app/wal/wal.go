package wal

import (
	"bufio"
	"fmt"
	"kwdb/app/commands"
	"log"
	"os"
)

func Write(text string) {

	file, err := os.OpenFile("./data/wal1.txt", os.O_APPEND, 066)

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

func Backup(backup_chan chan) {
	file, err := os.OpenFile("./data/wal1.txt", os.O_APPEND, 066)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		backup_chan <- scanner.Text()
		if err != nil {
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
