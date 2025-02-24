package errorpkg

import "fmt"

var (
	ErrHashFunctionCompute       = fmt.Errorf("error hashing function")
	ErrHashFunctionIndexOutRange = fmt.Errorf("partition index out of range")
	ErrUnknownDriver             = fmt.Errorf("unknown driver")

	ErrorParse                  = fmt.Errorf("ошибка разбора строки")
	ErrorParseTTL               = fmt.Errorf("ошибка чтения ttl")
	ErrorParseParameterNotFound = fmt.Errorf("передан неизвестный параметр")

	ErrorTcpReadAnswer       = fmt.Errorf("ошибка разбора ответа сервера")
	ErrorTcpSendMessage      = fmt.Errorf("ошибка отправки сообщения серверу")
	ErrorTcpSetUpConnections = fmt.Errorf("ошибка установки соединения с сервером")

	ErrCmdKeyNotFound = fmt.Errorf("key not found")
)
