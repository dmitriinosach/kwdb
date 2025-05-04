package backup

import (
	"bufio"
	"context"
	"kwdb/app/storage"
	"kwdb/internal/helper/file_system"
	"kwdb/internal/helper/informer"
	"log"
	"os"
	"time"
)

const (
	backupPath = "./data/backup/wal1.txt"

	BackupEndCtx = iota
	BackupEndTime
)

var backupFile *os.File

// TODO: переделать, долен решать другую задачу
func Write(text []byte) {

	if backupFile == nil {
		backupFile, _ = file_system.ReadOrCreate(backupPath)
	}

	_, err := backupFile.Write(text)
	_, err = backupFile.Write([]byte{0x0A})

	if err != nil {
		informer.PrintCli("Ошибка записи бэкапа")
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
				res <- BackupEndCtx
				return
			default:
			}

			for {
				select {
				case rc <- scanner.Text():
					break
				case <-tl.C:
					res <- BackupEndTime
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
