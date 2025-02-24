package app

import (
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"os"
)

type ConfigEnv struct {
	HOST   string
	PORT   string
	DRIVER string
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

	return Config, nil
}
