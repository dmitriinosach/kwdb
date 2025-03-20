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

		ifc := commands.List[commands.CommandSet]

		args := new(commands.Arguments)

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

	helper.InfChan <- "http://" + app.Config.Host + ":" + app.Config.Port + " ожидает подключений"

	err := http.ListenAndServe(app.Config.Host+":713", nil)

	if err != nil {
		helper.InfChan <- "http://" + app.Config.Host + ":" + app.Config.Port + " прекратил работу: " + err.Error()
		return
	}

}
