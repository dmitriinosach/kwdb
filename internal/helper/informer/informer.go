package informer

import (
	"context"
	"fmt"
	cprntr "kwdb/internal/helper/color_printer"
	"kwdb/internal/helper/common"
	"kwdb/internal/helper/flogger"
)

var InfChan = make(chan string, 1)

func Run(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			cprntr.PrintRed(common.GetPrefixNow() + "informer остановлен")
			flogger.Flogger.WriteString("shutdown...")
			return
		case m := <-InfChan:
			m = common.GetPrefixNow() + m
			fmt.Println(m)
		}
	}
}
