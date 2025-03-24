package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"kwdb/app/errorpkg"
	"os"
	"strconv"
)

// ConfigEnv структура для хранения конфигурации приложения
type ConfigEnv struct {
	// Host ip\domain базы
	Host string

	// Port порт базы данных
	Port string

	// Driver строковое имя драйвера напр. hashmap
	Driver string

	// Partitions кол-во партиций внутри хранилища
	Partitions int

	// LogPath путь хранения файлов логирования
	LogPath string

	// MemLimit лимит памяти хранилища
	MemLimit uint64
}

var Config ConfigEnv

// InitConfigs загрузка конфигураций приложения из env файла
func InitConfigs() error {

	if err := godotenv.Load(); err != nil {
		return fmt.Errorf(errorpkg.ErrEnvLoad)
	}

	exist := false

	Config.Host, exist = os.LookupEnv("SERVER_HOST")
	if !exist {
		return fmt.Errorf(errorpkg.ErrEnvParameterMissed + "SERVER_HOST")
	}

	Config.Port, exist = os.LookupEnv("SERVER_PORT")
	if !exist {
		return fmt.Errorf(errorpkg.ErrEnvParameterMissed + "SERVER_PORT")
	}

	Config.Driver, exist = os.LookupEnv("DATABASE_DRIVER")
	if !exist {
		return fmt.Errorf(errorpkg.ErrEnvParameterMissed + "DATABASE_DRIVER")
	}

	parts := ""
	parts, exist = os.LookupEnv("DATABASE_DRIVER_PARTITIONS")
	Config.Partitions, _ = strconv.Atoi(parts)

	if !exist {
		return fmt.Errorf(errorpkg.ErrEnvParameterMissed + "DATABASE_DRIVER")
	}

	Config.LogPath, exist = os.LookupEnv("LOG_PATH")
	if !exist {
		return fmt.Errorf(errorpkg.ErrEnvParameterMissed + "LOG_PATH")
	}

	Config.MemLimit = 100

	return nil
}
