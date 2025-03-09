package main

import (
	"regexp"
	"strconv"
)

type Reply struct {
	Result       string
	ResultErrors string
}

// разбираем ответ приложения по формату {int длинна ответа}:{ответ}:{ошибки}
func parseReply(r []byte) *Reply {
	rep := new(Reply)

	bodyLenFind := regexp.MustCompile(`^(\d+):`)
	matches := bodyLenFind.FindAllString(string(r), -1)

	bodyStart := len(matches[0])
	bodyLen, _ := strconv.Atoi(matches[0][0 : bodyStart-1])

	bodyEnd := bodyStart + bodyLen

	rep.Result = string(r[bodyStart:bodyEnd])
	rep.ResultErrors = string(r[bodyEnd:])

	return rep
}
