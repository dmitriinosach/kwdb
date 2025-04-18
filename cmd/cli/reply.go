package main

import (
	"encoding/binary"
)

type Reply struct {
	Result string
	Errors string

	Raw []byte
}

func NewReply(raw []byte) *Reply {
	r := &Reply{
		Raw: raw,
	}

	r.parse()

	return r
}

// parse разбираем ответ приложения по формату {uint32 длинна ответа}:{ответ}:{ошибки}
func (r *Reply) parse() {
	mySlice := r.Raw[0:4]

	l := int(binary.BigEndian.Uint32(mySlice))

	bodyEnd := len(r.Raw) - l

	r.Result = string(r.Raw[4 : bodyEnd+1])

	r.Errors = string(r.Raw[bodyEnd:])
}
