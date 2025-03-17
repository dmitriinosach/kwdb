package storage

import (
	"kwdb/app/errorpkg"
	"kwdb/app/storage/displacement"
	"kwdb/app/storage/driver"
	"kwdb/app/storage/driver/mapstd"
	"kwdb/app/storage/driver/syncmap"
	"sync"
)

var (
	Storage driver.Driver
	once    sync.Once
)

func Init(driverName string, partitionsCount int) (err error) {
	// TODO: флагами получить интерфес драйверов
	//TODO лишнее
	once.Do(func() {
		switch driverName {
		case mapstd.DriverName:

			Storage = mapstd.NewHashMapStandard(partitionsCount)
		case syncmap.DriverName:
			Storage = syncmap.NewSyncMap(partitionsCount)
		default:
			err = errorpkg.ErrUnknownDriver
		}

		setMemPolicy()
	})

	return
}

func setMemPolicy() {
	var policy displacement.Policy

	policy = displacement.NewLRU()

	Storage.SetMemPolicy(policy)
}
