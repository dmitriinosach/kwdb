package main

import (
	"fmt"
	"kwdb/app/errorpkg"
	"net"
	"time"
)

func goe(pac int) {

	for i := pac * 1000; i < (pac*1000)+1000; i++ {
		//send("SET value=cacheafgljgfjkgfjklgfdsjkgdfkjlgdfsljkgfdsljkgfdljkfgdsljkgfdsljkgfdsljk" + strconv.Itoa(i) + " key=" + strconv.Itoa(i))
	}
}

func send(message string) (string, error) {

	conn, err := net.Dial("tcp", cliConfig.connectionHost+":"+cliConfig.connectionPort)

	if err != nil {
		return "", errorpkg.ErrorTcpSetUpConnections
	}

	defer conn.Close()

	err = conn.SetReadDeadline(time.Now().Add(time.Second * 1))
	err = conn.SetWriteDeadline(time.Now().Add(time.Second * 1))

	if err != nil {
		return "", err
	}

	// отправляем сообщение серверу
	if n, err := conn.Write([]byte(":" + message)); n == 0 || err != nil {
		return "", errorpkg.ErrorTcpSendMessage
	}

	// получем ответ

	buff := make([]byte, 1024)
	n, err := conn.Read(buff)

	if err != nil {
		return "", errorpkg.ErrorTcpReadAnswer
	}

	response := buff[:n]

	res := parseReply(response)

	if len(res.ResultErrors) > 0 {
		return "", fmt.Errorf(res.ResultErrors)
	}

	return res.Result, nil
}
