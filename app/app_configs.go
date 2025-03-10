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
	HOST       string
	PORT       string
	DRIVER     string
	PARTITIONS int
	LogPath    string
}

var Config ConfigEnv

type SysInfo struct {
	Started time.Time
}

var SysInfoData SysInfo

func InitConfigs() (ConfigEnv, error) {

	if err := godotenv.Load(); err != nil {
		return ConfigEnv{}, errors.Wrap(err, errorpkg.ErrEnvLoad)
	}

	exist := false

	Config.HOST, exist = os.LookupEnv("SERVER_HOST")
	if !exist {
		return Config, errors.New(errorpkg.ErrEnvParameterMissed + "SERVER_HOST")
	}

	Config.PORT, exist = os.LookupEnv("SERVER_PORT")
	if !exist {
		return Config, errors.New(errorpkg.ErrEnvParameterMissed + "SERVER_PORT")
	}

	Config.DRIVER, exist = os.LookupEnv("DATABASE_DRIVER")
	if !exist {
		return Config, errors.New(errorpkg.ErrEnvParameterMissed + "DATABASE_DRIVER")
	}

	partitions := ""
	partitions, exist = os.LookupEnv("DATABASE_DRIVER_PARTITIONS")
	Config.PARTITIONS, _ = strconv.Atoi(partitions)

	if !exist {
		return Config, errors.New(errorpkg.ErrEnvParameterMissed + "DATABASE_DRIVER")
	}

	Config.LogPath, exist = os.LookupEnv("LOG_PATH")
	if !exist {
		return Config, errors.New(errorpkg.ErrEnvParameterMissed + "LOG_PATH")
	}

	return Config, nil
}
