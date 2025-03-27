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

func Serve() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		ctx := context.Background()
		cs := "SET key=" + r.URL.Query().Get("key") + " value=" + r.URL.Query().Get("value")

		res, err := commands.SetAndRun(ctx, cs)

		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}

		fmt.Fprintf(w, res)
	})

	informer.InfChan <- "http://" + app.Config.Host + ":" + app.Config.Port + " ожидает подключений"

	err := http.ListenAndServe(app.Config.Host+":713", nil)

	if err != nil {
		informer.InfChan <- "http://" + app.Config.Host + ":" + app.Config.Port + " прекратил работу: " + err.Error()
		return
	}

}
