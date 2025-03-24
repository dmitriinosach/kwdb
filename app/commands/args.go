// Copyright 2025 @dmitrii_nosach

package commands

import (
	"fmt"
	"github.com/google/shlex"
	"kwdb/app/errorpkg"
	"strconv"
	"strings"
)

// separator символ разделителя аргумента и значения в строке
const separator string = "="

type arguments struct {
	//Имя команды для определения
	CmdName string

	//аргументы команды
	Key   string
	Value string
	TTL   int
}

func newArgsFromString(s string) (*arguments, error) {

	args := new(arguments)

	parsedLine, err := shlex.Split(s)

	if err != nil {
		return args, errorpkg.ErrorParse
	}

	for number, tag := range parsedLine {
		if number == 0 {
			_, nameError := withCmdName(args, tag)
			if nameError != nil {
				return nil, nameError
			}
		} else {
			split := strings.Split(tag, separator)

			switch split[0] {
			case "key", "k":
				args.Key = split[1]
			case "value", "v":
				args.Value = split[1]
			case "ttl", "t":
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

func withCmdName(args *arguments, tag string) (*arguments, error) {
	cmdMaxLen := 10

	if len(tag) > cmdMaxLen {
		return nil, fmt.Errorf("name too long")
	}

	args.CmdName = tag

	return args, nil
}
