package http

import (
	"context"
	"fmt"
	"kwdb/app"
	"kwdb/app/commands"
	"net/http"
)

var Server *srv

// TODO: семафоры

func Serve(ctx context.Context) {
	Server = NewServer()
	// мидвалвары и семафор
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		cs := "SET key=" + r.URL.Query().Get("key") + " value=" + r.URL.Query().Get("value")

		res, err := commands.SetAndRun(cs)

		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}

		fmt.Fprintf(w, res)
	})

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
