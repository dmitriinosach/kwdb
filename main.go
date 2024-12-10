package main

import (
	"fmt"
	"kwdb/app"
	"kwdb/app/api"
)

func main() {

	go api.HandleCLI()

	fmt.Printf("Приложение запущено\n")

	app.Serve()

}
