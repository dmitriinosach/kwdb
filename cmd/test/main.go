package main

import (
	"fmt"
	"runtime"
)

// просто для теста
func main() {
	initialized := runtime.NumGoroutine()

	fmt.Println(initialized)
}
