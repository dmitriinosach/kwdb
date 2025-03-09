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

		ifc := commands.List[commands.CommandSet]

		args := new(commands.CommandArguments)

		args.CmdName = "SET"
		args.Key = r.URL.Query().Get("key")
		args.Value = r.URL.Query().Get("value")
		args.TTL = 0

		fmt.Printf("%v", args)

		ctx := context.Background()

		ifc.SetArgs(ctx, args)

		res, _ := ifc.Execute(ctx)
		fmt.Fprintf(w, res)

	})

	helper.InfChan <- "http://" + app.Config.HOST + ":" + app.Config.PORT + " ожидает подключений"

	err := http.ListenAndServe(app.Config.HOST+":713", nil)

	if err != nil {
		helper.InfChan <- "http://" + app.Config.HOST + ":" + app.Config.PORT + " прекратил работу: " + err.Error()
		return
	}

}
