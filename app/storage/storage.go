package storage

import (
	"kwdb/app/errorpkg"
	"kwdb/app/storage/cell"
	"kwdb/app/storage/displacement"
	"kwdb/app/storage/driver/mapstd"
	"kwdb/app/storage/driver/syncmap"
	"sync"
	"time"
)

var (
	Storage     Driver
	once        sync.Once
	Status      = new(status)
	cleanerChan chan string
	Lru         displacement.Policy
)

type Driver interface {
	Get(key string) (*cell.Cell, error)
	Set(key string, value string, ttl int) error
	Delete(key string) error
	Info() string
	GetVaultMap() map[string]*cell.Cell
	Truncate() bool
	Cleaner(cc chan string)
}

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
