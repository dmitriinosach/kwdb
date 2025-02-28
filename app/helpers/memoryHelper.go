package helpers

import (
	"fmt"
	"kwdb/app"
	"runtime"
)

func printAlloc() {
	var stat runtime.MemStats

	runtime.ReadMemStats(&stat)

	app.InfChan <- fmt.Sprintf("Alloc = %v MiB", stat.Alloc/1024/1024)

}
