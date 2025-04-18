package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// просто для теста
func main() {

	s := []byte("hello world sdf qwer fg ;iajsdh fpoashdfpoiahsdofwe r wr йц")
	l := uint32(len(s))
	fmt.Println(l)
	var bl = make([]byte, 4)
	buf := bytes.NewBuffer([]byte{})
	// Little-endian порядок байтов
	binary.BigEndian.PutUint32(bl[:], l)
	fmt.Println(bl)
	buf.Write(bl)
	buf.Write(s)

	fmt.Println(buf.String())
	fmt.Println(buf.Bytes())

	res := buf.Bytes()
	var mySlice = res[0:4]
	fmt.Printf("mySlice:%v\n", mySlice)
	lr := binary.BigEndian.Uint32(mySlice)

	fmt.Println(lr)

}
