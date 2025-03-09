package main

import (
	"fmt"
	"kwdb/app/errorpkg"
	"net"
	"regexp"
	"strconv"
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

	//GET
	// отправляем сообщение серверу
	if n, err := conn.Write([]byte(":" + message)); n == 0 || err != nil {
		return "", errorpkg.ErrorTcpSendMessage
	}

	// получем ответ

	buff := make([]byte, 1024)
	n, err := conn.Read(buff)

	response := buff[:n]
	if err != nil {
		return "", errorpkg.ErrorTcpReadAnswer
	}

	respLen := regexp.MustCompile(`^sign:(\d+):`)
	matches := respLen.FindAllString(string(response), -1)
	bodyStart := len(matches[0])

	bodyLen, _ := strconv.Atoi(matches[0][5 : bodyStart-1])

	bodyEnd := bodyStart + bodyLen
	responseBody := response[bodyStart:bodyEnd]
	errorBody := response[bodyEnd+1:]

	conn.Close()

	if len(errorBody) > 0 {
		return "", fmt.Errorf(string(errorBody))
	}

	return string(responseBody), nil
}
