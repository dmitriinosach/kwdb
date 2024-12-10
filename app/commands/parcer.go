package commands

import (
	"fmt"
	"github.com/google/shlex"
	"strconv"
	"strings"
)

func Parce(message string) CommandArguments {

	parsedLine, _ := shlex.Split(message)

	args := CommandArguments{}

	for number, tag := range parsedLine {
		if number == 0 {
			args.Name = tag
		} else {
			parameter := strings.Split(tag, "=")

			switch parameter[0] {
			case "key":
				args.Key = parameter[1]
			case "value":
				args.Value = parameter[1]
			case "ttl":
				args.TTL, _ = strconv.Atoi(parameter[1])
			default:
				fmt.Println("Unknown parameter:", parameter[0])
			}
		}
	}

	return args
}
