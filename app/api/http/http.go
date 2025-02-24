package http

import (
	"context"
	"fmt"
	"kwdb/app/commands"
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

	err := http.ListenAndServe("localhost:713", nil)
	if err != nil {
		return
	}
}
