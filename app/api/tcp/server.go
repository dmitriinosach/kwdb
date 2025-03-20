package tcp

import (
	"context"
	"kwdb/app"
	"kwdb/app/api"
	"kwdb/internal/helper"
	"kwdb/internal/helper/logger"
	"net"
	"os"
	"os/signal"
	"strconv"
)

func Serve() {

	listen, err := net.Listen("tcp", app.Config.Host+":"+app.Config.Port)

	ctx := context.Background()

	signal.NotifyContext(ctx, os.Interrupt)

	if err != nil {
		os.Exit(1)
	}

	defer listen.Close()

	helper.InfChan <- "tcp://" + app.Config.Host + ":" + app.Config.Port + " ожидает подключений"

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		conn, err := listen.Accept()
		if err != nil {
			logger.Write(err.Error(), "")
			continue
		}

		go tpcHandle(ctx, conn)
	}
}

func tpcHandle(ctx context.Context, conn net.Conn) {

	var result string
	buffer := make([]byte, 1024)

	bufferLength, err := conn.Read(buffer)

	if err != nil {
		r := ""
		//TODO: всрато выглядит
		answer(makeReply(&r, err), conn)
		logger.Write(err.Error(), "")

		return
	}

	message := string(buffer[1:bufferLength])

	//TODO:избавится от msg и парсера, перейти на структуру и байты
	result, err = api.ExecMsg(ctx, message)
	//TODO: всрато выглядит
	answer(makeReply(&result, err), conn)

	defer conn.Close()
}

// TODO: нужа или нет передача по ссылке
func makeReply(r *string, e error) []byte {

	reply := strconv.Itoa(len(*r)) + ":" + *r

	if e != nil {
		reply += (e).Error()
	}

	return []byte(reply)
}

func answer(r []byte, conn net.Conn) {
	_, err := conn.Write(r)
	if err != nil {
		logger.Write("Не удалось ответить серверу:"+err.Error(), "")
		return
	}
}
