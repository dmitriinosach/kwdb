package http

import (
	"context"
	"fmt"
	"kwdb/app"
	"kwdb/app/commands"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
)

var Server *srv

// TODO: семафоры

func Serve(ctx context.Context) {
	Server = NewServer()
	// мидвалвары и семафор
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		cs := "SET key=" + r.URL.Query().Get("key") + " value=" + r.URL.Query().Get("value")

		res, err := commands.SetAndRun([]byte(cs))

		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}

		fmt.Fprintf(w, string(res))
	})

	handler.HandleFunc("/debug/pprof/", pprof.Index)
	handler.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	handler.HandleFunc("/debug/pprof/profile", pprof.Profile)
	handler.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	handler.HandleFunc("/debug/pprof/trace", pprof.Trace)
	handler.Handle("/debug/pprof/block", pprof.Handler("block"))
	handler.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	handler.Handle("/debug/pprof/mutex", pprof.Handler("mutex"))
	handler.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	handler.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))

	Server.handler = handler
	Server.server = &http.Server{
		Addr:    Server.config.soc,
		Handler: handler,
	}

	app.WithHttp(Server.server)

	app.InfChan <- "http://" + Server.config.soc + " ожидает подключений"

	if err := Server.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		app.InfChan <- "http://" + Server.config.soc + " прекратил работу: " + err.Error()
	}
}
