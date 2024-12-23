package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type ConfigEnv struct {
	HOST   string
	PORT   string
	DRIVER string
}

var Config ConfigEnv

func InitConfigs() (ConfigEnv, error) {

	if err := godotenv.Load(); err != nil {
		log.Print("Файл с переменными окружения не найден")
	}

	host, exists := os.LookupEnv("SERVER_HOST")
	if !exists {
		return Config, fmt.Errorf("не установлен хост для приложения")
	}
	Config.HOST = host

	port, exists := os.LookupEnv("SERVER_PORT")
	if !exists {
		return Config, fmt.Errorf("не установлен порт для приложения")
	}
	Config.PORT = port

	driver, exists := os.LookupEnv("DATABASE_DRIVER")
	if !exists {
		return Config, fmt.Errorf("не установлен драйвер в настройках")
	}
	Config.DRIVER = driver

	return Config, nil
}
