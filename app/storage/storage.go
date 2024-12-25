package storage

import (
	"sync"
	"time"

	"kwdb/app/storage/driver"
)

var dict = map[string]driver.Interface{
	"hash": &driver.HashMapStandard{},
}

var (
	Storage driver.Interface
	once    sync.Once
)

var Started = time.Now()

func Init(dbDriver string) (err error) {
	once.Do(func() {
		hash, ok := dict[dbDriver]
		if ok {
			Storage = hash
			Storage.Init()
		}

		if Storage == nil {
			err = driver.ErrUnknownDriver
		}
	})

	return
}
