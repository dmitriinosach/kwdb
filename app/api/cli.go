package api

import (
	"bufio"
	"context"
	"fmt"
	"kwdb/app"
	"os"
)

func HandleCLI() {
	reader := bufio.NewReader(os.Stdin)
	ctx := context.Background()
	for {
		fmt.Print("kwdb-cli: ")
		message, _ := reader.ReadString('\n')

		result, err := app.HandleMessage(ctx, message)

		if err != nil {
			fmt.Println("Ошибка выполнения запроса", err)
		}

		fmt.Println(result)
	}
}
