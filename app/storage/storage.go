package storage

import (
	"kwdb/app/storage/cell"
	"kwdb/app/storage/displacement"
	"kwdb/app/storage/driver"
	"time"
)

var (
	Storage     Driver
	Status      = new(status)
	cleanerChan = make(chan string)
	Lru         displacement.Policy
)

type Driver interface {
	Get(key string) (*cell.Cell, error)
	Set(key string, value []byte, ttl int) error
	Delete(key string) error
	Info() []byte
	Truncate() bool
	Cleaner(cc chan string)
	Flush() error
}

func Init(driverName string, partitionsCount int) (err error) {
	// TODO: флагами получить интерфес драйверов
	//TODO лишнее

	mapType := "std"
	Lru = displacement.NewLRU(cleanerChan)

	Storage = driver.NewHashMapStandard(
		partitionsCount,
		Lru,
		mapType,
	)

	Status.Started = time.Now()
	Status.DriverName = driverName

	go displacement.RunWatcher(Lru)

	go Storage.Cleaner(cleanerChan)

	return
}
