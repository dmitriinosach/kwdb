package main

import (
	"fmt"
	"runtime"
)

func main() {
	initialized := runtime.NumGoroutine()

	fmt.Println(initialized)
}
