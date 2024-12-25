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

	host, exists := os.LookupEnv("SERVER_HOST")
	if !exists {
		return Config, errors.New(EnvParameterMissed + "SERVER_HOST")
	}
	Config.HOST = host

	port, exists := os.LookupEnv("SERVER_PORT")
	if !exists {
		return Config, errors.New(EnvParameterMissed + "SERVER_PORT")
	}
	Config.PORT = port

	driver, exists := os.LookupEnv("DATABASE_DRIVER")
	if !exists {
		return Config, errors.New(EnvParameterMissed + "DATABASE_DRIVER")
	}

	Config.DRIVER = driver

	return Config, nil
}
