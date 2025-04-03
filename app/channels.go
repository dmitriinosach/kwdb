package app

import (
	"context"
	cprntr "kwdb/internal/helper/color_printer"
	"kwdb/internal/helper/informer"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var InfChan = make(chan string)
var sigs = make(chan os.Signal, 1)
var tcpListener *net.Listener
var httpListener *http.Server

func ChanHandler(ctx context.Context, wg *sync.WaitGroup, cancel context.CancelFunc) {

	wg.Add(1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case m := <-InfChan:
			informer.PrintCli(m)
		case <-ctx.Done():
			close(sigs)
			close(InfChan)
			if httpListener != nil {
				httpListener.Close()
			}
			wg.Done()
			return
		case sig := <-sigs:
			if sig == syscall.SIGTERM || sig == syscall.SIGINT {
				cancel()
				cprntr.PrintRed("Выключение")
			}
		}
	}
}

func WithTcp(listen *net.Listener) {
	tcpListener = listen
}

func WithHttp(listen *http.Server) {
	httpListener = listen
}
