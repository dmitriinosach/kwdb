// Copyright 2025 @dmitrii_nosach

package commands

import (
	"errors"
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
	Value []byte
	TTL   int
}

func newArgsFromString(s string) (*arguments, error) {

	args := &arguments{}

	parsedLine, err := shlex.Split(s)

	if err != nil {
		return args, errorpkg.ErrorParse
	}

	for number, tag := range parsedLine {
		if number == 0 {
			nameError := args.setCmdName(tag)
			if nameError != nil {
				return nil, nameError
			}
		} else {
			split := strings.Split(tag, separator)

			switch split[0] {
			case "key", "k":
				args.Key = split[1]
			case "value", "v":
				args.Value = []byte(split[1])
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

func (a *arguments) setCmdName(tag string) error {
	cmdMaxLen := 10

	if len(tag) > cmdMaxLen {
		return errors.New("")
	}

	a.CmdName = tag

	return nil
}
