package tcp

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"kwdb/app"
	"kwdb/app/commands"
	"kwdb/internal/helper/flogger"
	"net"
	"os"
	"runtime"
)

var sem = make(chan interface{}, runtime.NumCPU())

func Serve(ctx context.Context) {

	addr := net.JoinHostPort(app.Config.Host, app.Config.Port)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app.WithTcp(&listen)

	defer listen.Close()

	app.InfChan <- "tcp://" + app.Config.Get("Host").(string) + ":" + app.Config.Port + " ожидает подключений"

	for {
		select {
		case <-ctx.Done():
			return
		case sem <- 1:
			conn, err := listen.Accept()
			if err != nil {
				flogger.Flogger.WriteString(err.Error())
				continue
			}
			go handle(conn)
		default:
		}
	}
}

func handle(conn net.Conn) {

	defer conn.Close()

	buffer := make([]byte, 1024)

	bufferLength, err := conn.Read(buffer)

	if err != nil {

		reply([]byte{}, err, conn)

		flogger.Flogger.WriteString(err.Error())

		return
	}

	result, err := commands.SetAndRun(buffer[1:bufferLength])

	reply(result, err, conn)

	<-sem
}

// TODO: нужна или нет передача по ссылке &r
func reply(r []byte, e error, conn net.Conn) {

	buf := bytes.NewBuffer([]byte{})
	bl := make([]byte, 4)
	binary.BigEndian.PutUint32(bl[:], uint32(len(r)))
	buf.Write(bl)
	buf.Write(r)
	if e != nil {
		buf.Write([]byte((e).Error()))
	}

	_, err := conn.Write(buf.Bytes())

	if err != nil {
		flogger.Flogger.WriteString("Не удалось ответить серверу:" + err.Error())
	}
}
