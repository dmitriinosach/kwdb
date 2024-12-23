package api

import (
	"bufio"
	"fmt"
	"kwdb/app"
	"os"
)

func HandleCLI() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("kwdb-cli: ")
		message, _ := reader.ReadString('\n')

		result, err := app.HandleMessage(message)

		if err != nil {
			fmt.Println("Ошибка выполнения запроса", err)
		}

		fmt.Println(result)
	}
}
