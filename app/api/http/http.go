package http

import (
	"context"
	"fmt"
	"kwdb/app"
	"kwdb/app/commands"
	"kwdb/internal/helper"
	"net/http"
)

type httpHandler struct {
}

func Serve() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		ctx := context.Background()
		cs := "SET k=" + r.URL.Query().Get("key") + " v=" + r.URL.Query().Get("value")

		res, _ := commands.SetAndRun(ctx, cs)

		fmt.Fprintf(w, res)
	})

	helper.InfChan <- "http://" + app.Config.Host + ":" + app.Config.Port + " ожидает подключений"

	err := http.ListenAndServe(app.Config.Host+":713", nil)

	if err != nil {
		helper.InfChan <- "http://" + app.Config.Host + ":" + app.Config.Port + " прекратил работу: " + err.Error()
		return
	}

}
