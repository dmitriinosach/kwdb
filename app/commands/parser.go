package commands

import (
	"github.com/google/shlex"
	"kwdb/app/errorpkg"
	"strconv"
	"strings"
)

func parse(message string) (*CommandArguments, error) {
	args := new(CommandArguments)

	parsedLine, err := shlex.Split(message)

	if err != nil {
		return args, errorpkg.ErrorParse
	}

	for number, tag := range parsedLine {
		if number == 0 {
			cmdMaxLen := min(len(tag), 10)
			args.CmdName = tag[:cmdMaxLen]
		} else {
			parameter := strings.Split(tag, "=")

			switch parameter[0] {
			case "key", "k":
				args.Key = parameter[1]
			case "value", "v":
				args.Value = parameter[1]
			case "ttl":
				ttl, ok := strconv.Atoi(parameter[1])
				if ok != nil {
					return args, errorpkg.ErrorParseTTL
				}
				args.TTL = ttl
			default:
				return args, errorpkg.ErrorParseParameterNotFound
			}
		}
	}

	return args, nil
}
