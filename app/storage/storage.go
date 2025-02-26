package storage

import (
	"kwdb/app/errorpkg"
	"kwdb/app/storage/driver"
	"kwdb/app/storage/driver/mapstd"
	"kwdb/app/storage/driver/syncmap"
	"sync"
	"time"
)

var (
	Started = time.Now()
	Storage driver.Driver
	once    sync.Once
)

func Init(driverName string, partitionsCount int) (err error) {
	// TODO: флагами получить интерфес драйверов
	once.Do(func() {
		switch driverName {
		case mapstd.DriverName:
			Storage = mapstd.NewHashMapStandard(partitionsCount)
		case syncmap.DriverName:
			Storage = syncmap.NewSyncMap(partitionsCount)
		default:
			err = errorpkg.ErrUnknownDriver
		}

	})

	return
}
