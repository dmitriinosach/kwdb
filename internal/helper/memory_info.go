package helper

import (
	"fmt"
	"kwdb/internal/helper/informer"
	"runtime"
)

func printAlloc() {
	var stat runtime.MemStats

	runtime.ReadMemStats(&stat)

	informer.InfChan <- fmt.Sprintf("Alloc = %v MiB", stat.Alloc/1024/1024)

}

func MemStatInfo() string {
	var stat runtime.MemStats

	runtime.ReadMemStats(&stat)

	return fmt.Sprintf("alloc = %v MiB", stat.Alloc/1024/1024)
}

func AllocMB() uint64 {
	var stat runtime.MemStats

	runtime.ReadMemStats(&stat)

	return stat.Alloc / 1024 / 1024
}
