package app

import (
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"kwdb/app/errorpkg"
	"os"
	"strconv"
	"time"
)

type ConfigEnv struct {
	Host       string
	Port       string
	Driver     string
	Partitions int
	LogPath    string
	MemLimit   uint64
}

var Config ConfigEnv

type SysInfo struct {
	Started time.Time
}

var SysInfoData SysInfo

func InitConfigs() (ConfigEnv, error) {

	// рефакторинг

	if err := godotenv.Load(); err != nil {
		return ConfigEnv{}, errors.Wrap(err, errorpkg.ErrEnvLoad)
	}

	exist := false

	Config.Host, exist = os.LookupEnv("SERVER_HOST")
	if !exist {
		return Config, errors.New(errorpkg.ErrEnvParameterMissed + "SERVER_HOST")
	}

	Config.Port, exist = os.LookupEnv("SERVER_PORT")
	if !exist {
		return Config, errors.New(errorpkg.ErrEnvParameterMissed + "SERVER_PORT")
	}

	Config.Driver, exist = os.LookupEnv("DATABASE_DRIVER")
	if !exist {
		return Config, errors.New(errorpkg.ErrEnvParameterMissed + "DATABASE_DRIVER")
	}

	parts := ""
	parts, exist = os.LookupEnv("DATABASE_DRIVER_PARTITIONS")
	Config.Partitions, _ = strconv.Atoi(parts)

	if !exist {
		return Config, errors.New(errorpkg.ErrEnvParameterMissed + "DATABASE_DRIVER")
	}

	Config.LogPath, exist = os.LookupEnv("LOG_PATH")
	if !exist {
		return Config, errors.New(errorpkg.ErrEnvParameterMissed + "LOG_PATH")
	}

	Config.MemLimit = 100

	return Config, nil
}
