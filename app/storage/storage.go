package storage

import (
	"context"
	"kwdb/app/errorpkg"
	"kwdb/app/storage/displacement"
	"kwdb/app/storage/driver"
	"kwdb/app/storage/driver/mapstd"
	"kwdb/app/storage/driver/syncmap"
	"sync"
	"time"
)

var (
	Storage     driver.Driver
	once        sync.Once
	CleanerChan chan string
)

var Status = new(status)

func Init(driverName string, partitionsCount int) (err error) {
	// TODO: флагами получить интерфес драйверов
	//TODO лишнее
	once.Do(func() {
		switch driverName {
		case mapstd.DriverName:
			//сделать не экспортируемым
			Storage = mapstd.NewHashMapStandard(partitionsCount)
		case syncmap.DriverName:
			Storage = syncmap.NewSyncMap(partitionsCount)
		default:
			err = errorpkg.ErrUnknownDriver
			return
		}

		Status.Started = time.Now()
		Status.DriverName = driverName

		CleanerChan = make(chan string)

		go displacement.RunWatcher(displacement.NewLRU(CleanerChan))

		// потом отрефачить получателя клинера
		go func() {
			ctx := context.Background()
			for key := range CleanerChan {
				Storage.Delete(ctx, key)
			}
		}()
	})

	return
}
