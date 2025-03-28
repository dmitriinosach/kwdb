package main

import (
	"bytes"
	"kwdb/app/errorpkg"
	"net"
	"time"
)

func goe(pac int) {

	for i := pac * 1000; i < (pac*1000)+1000; i++ {
		//send("SET value=cacheafgljgfjkgfjklgfdsjkgdfkjlgdfsljkgfdsljkgfdljkfgdsljkgfdsljkgfdsljk" + strconv.Itoa(i) + " key=" + strconv.Itoa(i))
	}
}

func handle(message string) string {
	message = string(bytes.Trim([]byte(message), "\x00"))
	result, errors := send(message)

	if errors != nil {
		return errors.Error()
	}

	if len(result.Errors) > 1 {
		return result.Errors
	}

	return result.Result
}

func send(message string) (*Reply, error) {

	conn, err := net.Dial("tcp", cliConfig.connectionHost+":"+cliConfig.connectionPort)

	if err != nil {
		return nil, errorpkg.ErrorTcpSetUpConnections
	}

	defer conn.Close()

	err = conn.SetReadDeadline(time.Now().Add(time.Second * 1))
	err = conn.SetWriteDeadline(time.Now().Add(time.Second * 1))

	if err != nil {
		return nil, err
	}

	// отправляем сообщение серверу
	if n, err := conn.Write([]byte(":" + message)); n == 0 || err != nil {
		return nil, err
	}

	// получем ответ
	buff := make([]byte, 1024)

	n, err := conn.Read(buff)

	if err != nil {
		return nil, errorpkg.ErrorTcpReadAnswer
	}

	res := NewReply(buff[:n])

	return res, nil
}
