package tcp

import (
	"context"
	"kwdb/app"
	"kwdb/app/handlers"
	"kwdb/pkg/helper"
	"kwdb/pkg/helper/logger"
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
			logger.Write(err.Error())
			continue
		}

		go tpcHandle(ctx, conn)
	}
}

func tpcHandle(ctx context.Context, conn net.Conn) {

	buffer := make([]byte, 1024)

	bufferLength, err := conn.Read(buffer)

	if err != nil {
		logger.Write(err.Error())
	}

	message := string(buffer[1:bufferLength])

	result, err := handlers.HandleMessage(ctx, message)

	if err != nil {
		logger.Write(err.Error())
	}

	_, err = conn.Write(makeReply(result, err))
	if err != nil {
		logger.Write(err.Error())
		return
	}

	err = conn.Close()
}

func makeReply(r string, e error) []byte {
	reply := "sign:" + strconv.Itoa(len(r)) + ":"
	reply += r + ":"

	if e != nil {
		logger.Write(e.Error())
		reply += e.Error()
	}

	return []byte(reply)
}
