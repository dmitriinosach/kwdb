package informer

import (
	"fmt"
	"time"
)

var InfChan = make(chan string, 1)

func Informer() {
	for m := range InfChan {
		m = "[" + time.Now().Format("2006-01-02 15:04:05") + "] " + m
		fmt.Println(m)
	}
}
