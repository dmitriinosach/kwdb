package main

import (
	"regexp"
	"strconv"
)

type Reply struct {
	Result string
	Errors string
	Raw    []byte
}

func NewReply(raw []byte) *Reply {
	r := &Reply{
		Raw: raw,
	}

	r.parse()

	return r
}

// разбираем ответ приложения по формату {int длинна ответа}:{ответ}:{ошибки}
func (r *Reply) parse() {
	bodyLenFind := regexp.MustCompile(`^\d+:`)

	matches := bodyLenFind.FindAllString(string(r.Raw), -1)

	bodyStart := len(matches[0])
	bodyLen, _ := strconv.Atoi(matches[0][0 : bodyStart-1])
	bodyEnd := bodyStart + bodyLen

	r.Result = string(r.Raw)[bodyStart:bodyEnd]

	r.Errors = string(r.Raw)[bodyEnd:]
}
