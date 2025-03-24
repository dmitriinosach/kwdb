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

func MemStatInfo() string {
	var stat runtime.MemStats

	runtime.ReadMemStats(&stat)

	return fmt.Sprintf("alloc = %v MiB", stat.Alloc/1024/1024)
}
