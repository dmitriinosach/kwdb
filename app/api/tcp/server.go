package tcp

import (
	"context"
	"fmt"
	"kwdb/app"
	"kwdb/app/commands"
	"kwdb/internal/helper/flogger"
	"kwdb/internal/helper/informer"
	"net"
	"os"
	"strconv"
)

func Serve(ctx context.Context) {

	addr := net.JoinHostPort(app.Config.Host, app.Config.Port)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer listen.Close()

	informer.InfChan <- "tcp://" + app.Config.Host + ":" + app.Config.Port + " ожидает подключений"

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("tcp shutdown:" + addr)
				listen.Close()
				return
			}
		}
	}(ctx)

	for {
		conn, err := listen.Accept()
		if err != nil {
			flogger.Flogger.WriteString(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {

	defer conn.Close()

	buffer := make([]byte, 1024)

	bufferLength, err := conn.Read(buffer)

	if err != nil {

		reply("", err, conn)

		flogger.Flogger.WriteString(err.Error())

		return
	}

	message := string(buffer[1:bufferLength])

	//TODO:избавится от msg и парсера, перейти на структуру и байты
	result, err := commands.SetAndRun(message)

	reply(result, err, conn)
}

// TODO: нужна или нет передача по ссылке &r
func reply(r string, e error, conn net.Conn) {

	r = strconv.Itoa(len(r)) + ":" + r

	if e != nil {
		r += (e).Error()
	}

	_, err := conn.Write([]byte(r))

	if err != nil {
		flogger.Flogger.WriteString("Не удалось ответить серверу:" + err.Error())
	}
}
