package helper

import (
	"fmt"
	"net"
)

func LocalIp() string {
	ip := ""
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	for _, addr := range addrs {
		// Проверяем, что это IP-адрес, а не loopback
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {

				ip = ipNet.IP.String()
			}
		}
	}

	return ip
}
