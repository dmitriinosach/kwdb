package storage

import (
	"kwdb/app/storage/driver"
	"time"
)

var List = []driver.Interface{
	&driver.HashMapStandard{},
}

var Storage driver.Interface
var Started = time.Now()

func Init(dbDriver string) {

	for _, db := range List {
		if db.GetDriver() == dbDriver {
			Storage = db
			Storage.Init()
			break
		}
	}

	if Storage == nil {
		panic("Драйвер базы данных не найден")
	}
}
