package api

import (
	"bufio"
	"fmt"
	"kwdb/app/commands"
	"os"
)

func HandleCLI() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("kwdb-cli: ")
		message, _ := reader.ReadString('\n')

		cmd, err := commands.SetupCommand(message)

		if err != nil {
			fmt.Println(err)
			continue
		}

		result, err := cmd.Execute()

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(result)
	}
}
