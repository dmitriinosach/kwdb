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
	cleanerChan = make(chan string)
	Lru         displacement.Policy
)

type Driver interface {
	Get(key string) (*cell.Cell, error)
	Set(key string, value []byte, ttl int) error
	Delete(key string) error
	Info() []byte
	GetVaultMap() map[string]*cell.Cell
	Truncate() bool
	Cleaner(cc chan string)
	Flush() error
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

		go displacement.RunWatcher(Lru)

		go Storage.Cleaner(cleanerChan)
	})

	return
}
