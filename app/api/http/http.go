package http

import (
	"context"
	"fmt"
	"kwdb/app"
	"kwdb/app/commands"
	"kwdb/pkg/helper"
	"net/http"
)

type httpHandler struct {
}

func Serve() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		ifc := commands.List[commands.CommandInfo]
		ctx := context.Background()
		res, _ := ifc.Execute(ctx)
		fmt.Fprintf(w, res)
	})

	helper.InfChan <- "http://" + app.Config.HOST + ":" + app.Config.PORT + " ожидает подключений"

	err := http.ListenAndServe("localhost:713", nil)

	if err != nil {
		helper.InfChan <- "http://" + app.Config.HOST + ":" + app.Config.PORT + " прекратил работу: " + err.Error()
		return
	}

}
