package app

import (
	"context"
	"kwdb/app/handlers"
	"kwdb/app/logger"
	"net"
	"os"
	"os/signal"
	"strconv"
)

func ServeTCP() {

	listen, err := net.Listen("tcp", Config.HOST+":"+Config.PORT)
	ctx := context.Background()
	signal.NotifyContext(ctx, os.Interrupt)

	if err != nil {
		os.Exit(1)
	}

	defer listen.Close()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		conn, err := listen.Accept()
		if err != nil {
			logger.Write(err.Error())
			os.Exit(1)
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
	logger.Write(message)
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
