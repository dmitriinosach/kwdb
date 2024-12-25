package workers

import (
	"kwdb/app/storage"
	"kwdb/app/storage/driver"
	"time"
)

func CleanerRun() {

	vault := storage.Storage.GetVaultMap()

	for {
		cleaner(vault)
		time.Sleep(10 * time.Second)
	}
}

func cleaner(vault map[string]*driver.Cell) {
	var nowTime = time.Now()
	storage.Storage.Lock()
	count := 0
	for key, value := range vault {
		if value.TTL > 0 {
			cellTime := value.AddDate.Add(time.Duration(value.TTL))
			if cellTime.Before(nowTime) {
				storage.Storage.DeleteValue(key)
			}
		} else {

		}

		count++
		if count > 100 {
			break
		}
	}

	storage.Storage.Unlock()
}
