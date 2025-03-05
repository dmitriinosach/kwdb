package helper

import (
	"fmt"
	"runtime"
)

func printAlloc() {
	var stat runtime.MemStats

	runtime.ReadMemStats(&stat)

	InfChan <- fmt.Sprintf("Alloc = %v MiB", stat.Alloc/1024/1024)

}
