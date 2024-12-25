package commands

import (
	"fmt"
	"github.com/google/shlex"
	"strconv"
	"strings"
)

func Parse(message string) (*CommandArguments, error) {
	parsedLine, err := shlex.Split(message)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга строки: %s, ошибка: %w", message, err)
	}

	args := new(CommandArguments)

	for number, tag := range parsedLine {
		if number == 0 {
			cmdMaxLen := min(len(tag), 10)
			args.Name = tag[:cmdMaxLen]
		} else {
			parameter := strings.Split(tag, "=")

			switch parameter[0] {
			case "key":
				args.Key = parameter[1]
			case "value":
				args.Value = parameter[1]
			case "ttl":
				ttl, ok := strconv.Atoi(parameter[1])
				if ok != nil {
					return args, fmt.Errorf("ошибка чтения ttl: %v", parameter[1])
				}
				args.TTL = ttl
			default:
				return args, fmt.Errorf("неизвестный параметр: %v", parameter[0])
			}
		}
	}

	return args, nil
}
