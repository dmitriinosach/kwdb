package storage

import (
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
	Status      = new(status)
	cleanerChan chan string
)

var Lru displacement.Policy

func Init(driverName string, partitionsCount int) (err error) {
	// TODO: флагами получить интерфес драйверов
	//TODO лишнее
	once.Do(func() {
		
		Lru = displacement.NewLRU(cleanerChan)

		switch driverName {
		case mapstd.DriverName:
			//сделать не экспортируемым
			Storage = mapstd.NewHashMapStandard(partitionsCount, Lru)
		case syncmap.DriverName:
			Storage = syncmap.NewSyncMap(partitionsCount)
		default:
			err = errorpkg.ErrUnknownDriver
			return
		}

		Status.Started = time.Now()
		Status.DriverName = driverName

		cleanerChan = make(chan string)

		go displacement.RunWatcher(Lru)

		go Storage.Cleaner(cleanerChan)
	})

	return
}
