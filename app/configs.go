package app

import (
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"os"
	"strconv"
)

type ConfigEnv struct {
	HOST       string
	PORT       string
	DRIVER     string
	PARTITIONS int
	LogPath    string
}

var (
	EnvParameterMissed = "в настройках окружения не установленно: "
	EnvLoad            = "Ошибка инициализации env файла"
)

var Config ConfigEnv

func InitConfigs() (ConfigEnv, error) {

	if err := godotenv.Load(); err != nil {
		return ConfigEnv{}, errors.Wrap(err, EnvLoad)
	}

	exist := false

	Config.HOST, exist = os.LookupEnv("SERVER_HOST")
	if !exist {
		return Config, errors.New(EnvParameterMissed + "SERVER_HOST")
	}

	Config.PORT, exist = os.LookupEnv("SERVER_PORT")
	if !exist {
		return Config, errors.New(EnvParameterMissed + "SERVER_PORT")
	}

	Config.DRIVER, exist = os.LookupEnv("DATABASE_DRIVER")
	if !exist {
		return Config, errors.New(EnvParameterMissed + "DATABASE_DRIVER")
	}

	partitions := ""
	partitions, exist = os.LookupEnv("DATABASE_DRIVER_PARTITIONS")
	Config.PARTITIONS, _ = strconv.Atoi(partitions)

	if !exist {
		return Config, errors.New(EnvParameterMissed + "DATABASE_DRIVER")
	}

	Config.LogPath, exist = os.LookupEnv("LOG_PATH")
	if !exist {
		return Config, errors.New(EnvParameterMissed + "LOG_PATH")
	}

	return Config, nil
}
