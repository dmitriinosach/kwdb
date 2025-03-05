package tcp

import (
	"context"
	"kwdb/app"
	"kwdb/app/handlers"
	"kwdb/pkg/helper"
	"net"
	"os"
	"os/signal"
	"strconv"
)

func Serve() {

	listen, err := net.Listen("tcp", app.Config.HOST+":"+app.Config.PORT)

	ctx := context.Background()

	signal.NotifyContext(ctx, os.Interrupt)

	if err != nil {
		os.Exit(1)
	}

	defer listen.Close()

	helper.InfChan <- "tcp://" + app.Config.HOST + ":" + app.Config.PORT + " ожидает подключений"

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		conn, err := listen.Accept()
		if err != nil {
			helper.Write(err.Error())
			continue
		}

		go tpcHandle(ctx, conn)
		helper.InfChan <- "Соединение принято:" + conn.RemoteAddr().String()
	}
}

func tpcHandle(ctx context.Context, conn net.Conn) {

	buffer := make([]byte, 1024)

	bufferLength, err := conn.Read(buffer)

	if err != nil {
		helper.Write(err.Error())
	}

	message := string(buffer[1:bufferLength])

	result, err := handlers.HandleMessage(ctx, message)

	if err != nil {
		helper.Write(err.Error())
	}

	_, err = conn.Write(makeReply(result, err))
	if err != nil {
		helper.Write(err.Error())
		return
	}

	err = conn.Close()

	helper.InfChan <- "Соединение обработано и закрыто:" + conn.RemoteAddr().String()
}

func makeReply(r string, e error) []byte {
	reply := "sign:" + strconv.Itoa(len(r)) + ":"
	reply += r + ":"

	if e != nil {
		helper.Write(e.Error())
		reply += e.Error()
	}

	return []byte(reply)
}
