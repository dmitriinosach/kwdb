package informer

import (
	"context"
	"fmt"
	"kwdb/internal/helper/common"
)

func Run(ctx context.Context) {

}

func PrintCli(m string) {
	m = common.GetPrefixNow() + m
	fmt.Println(m)
}
