package http

import (
	"context"
	"fmt"
	"kwdb/app"
	"kwdb/app/commands"
	"kwdb/internal/helper/informer"
	"net/http"
)

type httpHandler struct {
}

// TODO: семафоры

func Serve(ctx context.Context) {

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

	srv := &http.Server{
		Addr:    app.Config.Host + ":713",
		Handler: handler,
	}

	go func() {
		<-ctx.Done()
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Println("http server Shutdown err:" + err.Error())
		}
		fmt.Println("http server Shutdown")
	}()

	informer.InfChan <- "http://" + app.Config.HttpHost + ":" + app.Config.HttpPort + " ожидает подключений"

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		informer.InfChan <- "http://" + app.Config.Host + ":" + app.Config.Port + " прекратил работу: " + err.Error()
	}
}
