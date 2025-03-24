package tcp

import (
	"context"
	"kwdb/app"
	"kwdb/app/commands"
	"kwdb/internal/helper"
	"kwdb/internal/helper/flogger"
	"net"
	"os"
	"os/signal"
	"strconv"
)

func Serve() {

	addr := net.JoinHostPort(app.Config.Host, app.Config.Port)

	listen, err := net.Listen("tcp", addr)

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
			flogger.Write(err.Error(), "")
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
		reply(r, err, conn)

		flogger.Write(err.Error(), "")

		return
	}

	message := string(buffer[1:bufferLength])

	//TODO:избавится от msg и парсера, перейти на структуру и байты
	result, err = commands.SetAndRun(ctx, message)

	reply(result, err, conn)

	conn.Close()
}

// TODO: нужна или нет передача по ссылке &r
func reply(r string, e error, conn net.Conn) {

	r = strconv.Itoa(len(r)) + ":" + r

	if e != nil {
		r += (e).Error()
	}

	_, err := conn.Write([]byte(r))

	if err != nil {
		flogger.Write("Не удалось ответить серверу:"+err.Error(), "")
	}
}
