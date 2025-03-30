package informer

import (
	"context"
	"fmt"
	"kwdb/internal/helper/flogger"
	"time"
)

var InfChan = make(chan string, 1)

func Informer(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			flogger.Flogger.WriteString("shutdown...")
			return
		case m := <-InfChan:
			m = "[" + time.Now().Format("2006-01-02 15:04:05") + "] " + m
			fmt.Println(m)
		}
	}
}
