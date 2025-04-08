package helper

import (
	"fmt"
	"kwdb/app"
	"runtime"
	"strconv"
)

func printAlloc() {
	var stat runtime.MemStats

	runtime.ReadMemStats(&stat)

	app.InfChan <- fmt.Sprintf("Alloc = %v MiB", stat.Alloc/1024/1024)

}

func MemStatInfo() string {
	var stat runtime.MemStats
	runtime.ReadMemStats(&stat)

	i := "Heap:" + strconv.Itoa(int(stat.HeapAlloc/1024/1024)) + " MB\n"
	i += "Stack:" + strconv.Itoa(int(stat.StackInuse/1024/1024)) + " MB\n"
	return i
}

func AllocMB() uint64 {
	var stat runtime.MemStats

	runtime.ReadMemStats(&stat)

	return stat.Alloc / 1024 / 1024
}
