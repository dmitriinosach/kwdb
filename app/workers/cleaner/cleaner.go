package cleaner

import (
	"fmt"
	"kwdb/app/storage"
	"kwdb/app/storage/driver"
	"kwdb/app/storage/driver/mapstd"
	"kwdb/app/storage/driver/syncmap"
	"kwdb/internal/helper"
)

func Run(partitions int) {

	switch storage.Storage.(type) {
	case *mapstd.HashMapStandard:
		/*workSem := make(chan struct{})
		var t = reflect.ValueOf(&storage.Storage).Type().Elem()
		fmt.Println(t)
		for i := 0; i < app.Config.PARTITIONS; i++ {
			go (storage.Storage).MakeTBCleaner(workSem)
		}*/
	case *syncmap.SyncMap:
		fmt.Println("f")

	default:
		fmt.Println("default")
	}

	helper.InfChan <- "Cleaner запущен"
}

func cleaner(vault map[string]*driver.Cell) {
	/*var nowTime = time.Now()
	storage.Storage.Lock()
	count := 0
	for key, value := range vault {
		if value.TTL > 0 {
			cellTime := value.AddDate.Add(time.Duration(value.TTL))
			if cellTime.Before(nowTime) {
				storage.Storage.Delete(key)
			}
		} else {

		}

		count++
		if count > 100 {
			break
		}
	}

	storage.Storage.Unlock()*/
}
