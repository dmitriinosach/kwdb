package app

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func ServeTCP() {

	listen, err := net.Listen("tcp", Config.HOST+":"+Config.PORT)

	if err != nil {
		os.Exit(1)
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		go tpcHandle(conn)
	}

}

func tpcHandle(conn net.Conn) {
	buffer := make([]byte, 1024)

	bufferLength, err := conn.Read(buffer)

	if err != nil {
		log.Println(err)
	}

	message := string(buffer[:bufferLength])

	result, err := HandleMessage(message)

	reply := "sign:" + strconv.Itoa(len(result)) + ":"
	reply += result + ":"
	if err != nil {
		reply += err.Error()
	}

	_, err = conn.Write([]byte(reply))
	if err != nil {
		return
	}

	err = conn.Close()
}
