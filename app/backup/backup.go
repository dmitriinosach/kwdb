package backup

import (
	"bufio"
	"context"
	"kwdb/app"
	"kwdb/app/storage"
	"kwdb/internal/helper/file_system"
	"log"
	"os"
	"time"
)

const backupPath = "./data/backup/wal1.txt"

const BACKUP_END_CTX = 1
const BACKUP_END_TIME = 2

var backupFile *os.File

// TODO: переделать, долен решать другую задачу
func Write(text string) {

	if backupFile == nil {
		backupFile, _ = file_system.ReadOrCreate(backupPath)
	}

	_, err := backupFile.WriteString(text + "\n")

	if err != nil {
		app.InfChan <- "Ошибка записи wal"
	}
}

func Backup(ctx context.Context) (<-chan string, chan int) {
	file, err := os.OpenFile("./data/backup/wal1.txt", os.O_APPEND, 066)
	if err != nil {
		log.Fatal(err)
	}

	rc := make(chan string)
	res := make(chan int, 1)
	tl := time.NewTimer(time.Second * 3)

	// TODO: переписать на ротацию журнала
	storage.Status.Restoring.Store(true)

	go func() {
		defer func() {
			defer close(rc)
			defer close(res)
			defer tl.Stop()
			defer file.Close()
			storage.Status.Restoring.Store(false)
		}()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			select {
			case <-ctx.Done():
				res <- BACKUP_END_CTX
				return
			default:
			}

			for {
				select {
				case rc <- scanner.Text():
					break
				case <-tl.C:
					res <- BACKUP_END_TIME
					return
				default:
					continue
				}

				break
			}

			tl.Reset(time.Second * 3)
		}
	}()

	return rc, res
}
