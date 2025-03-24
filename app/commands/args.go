package commands

import (
	"github.com/google/shlex"
	"kwdb/app/errorpkg"
	"strconv"
	"strings"
)

type arguments struct {
	CmdName string
	Key     string
	Value   string
	TTL     int
}

func NewArgsFromString(s string) (*arguments, error) {

	args := new(arguments)

	parsedLine, err := shlex.Split(s)

	if err != nil {
		return args, errorpkg.ErrorParse
	}

	for number, tag := range parsedLine {
		if number == 0 {
			args.CmdName, err = cmdName(tag)
			if err != nil {
				return args, err
			}
		} else {
			split := strings.Split(tag, "=")

			switch split[0] {
			case "key", "k":
				args.Key = split[1]
			case "value", "v":
				args.Value = split[1]
			case "ttl":
				ttl, ok := strconv.Atoi(split[1])
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

func cmdName(tag string) (string, error) {
	cmdMaxLen := 10

	if len(tag) > cmdMaxLen {
		return tag, errorpkg.ErrorParseCmdNotFound
	}

	return tag, nil
}
