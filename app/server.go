package app

import (
	"bytes"
	"fmt"
	"kwdb/app/commands"
	"kwdb/app/wal"
	"log"
	"net"
	"os"
	"sync"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

var someMapMutex = sync.RWMutex{}

func Serve() {
	listen, err := net.Listen(TYPE, HOST+":"+PORT)

	if err != nil {
		//log.Fatal(err)
		os.Exit(1)
	}
	// close listener
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		go handleRequest(conn)
	}

}

func handleRequest(conn net.Conn) {

	buffer := make([]byte, 0, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	message := string(bytes.TrimRight(buffer, "\x00, \n"))

	cmd, err := commands.SetupCommand(message)
	if err != nil || cmd == nil {
		fmt.Printf("command error: %v", err)

	} else {
		go execute(conn, cmd)
		go wal.Write(message)
	}
}

func execute(conn net.Conn, cmd commands.CommandInterface) {
	someMapMutex.Lock()
	result, _ := cmd.Execute()
	someMapMutex.Unlock()
	_, err := conn.Write([]byte(result))
	if err != nil {
		return
	}

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
}
