package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"kwdb/app/errorpkg"
	"os"
	"reflect"
	"strconv"
	"sync"
)

var Config = &ConfigEnv{}

// ConfigEnv структура для хранения конфигурации приложения
type ConfigEnv struct {
	// Host ip\domain базы
	Host string

	// Port порт базы данных
	Port string

	// HttpHost http domain базы
	HttpHost string

	// HttpPort порт базы данных
	HttpPort int

	// Driver строковое имя драйвера напр. hashmap
	Driver string

	// Partitions кол-во партиций внутри хранилища
	Partitions int

	// LogPath путь хранения файлов логирования
	LogPath string

	// MemLimit лимит памяти хранилища
	MemLimit uint64
}

func (env *ConfigEnv) Get(key string) interface{} {
	v := reflect.ValueOf(env).Elem()
	return v.FieldByName(key).Interface()
}

// InitConfigs загрузка конфигураций приложения из env файла
func InitConfigs() error {
	var ierr error
	var once sync.Once

	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			ierr = err
			return
		}

		// Получаем reflect значение структуры
		v := reflect.ValueOf(Config).Elem()

		// Проходим по всем полям
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			fieldName := v.Type().Field(i).Name

			if field.CanSet() {
				cnf, me := configMatches[fieldName]
				if cnf.key == "" {
					ierr = fmt.Errorf("мэтч не содержит ключа env: " + fieldName)
					return
				}

				if !me {
					ierr = fmt.Errorf("не установлен мэтчи настройки: " + fieldName)
					return
				}

				cv, exist := os.LookupEnv(cnf.key)
				if !exist && cnf.defVal == "" {
					ierr = fmt.Errorf(errorpkg.ErrEnvParameterMissed + fieldName)
					return
				}

				if cv == "" {
					cv = cnf.defVal
				}

				switch field.Kind() {
				case reflect.String:
					field.SetString(cv)
				case reflect.Int:
					in, _ := strconv.Atoi(cv)
					field.SetInt(int64(in))
				case reflect.Uint64:
					in, _ := strconv.Atoi(cv)
					field.SetUint(uint64(in))
				default:
					ierr = fmt.Errorf("не обработанный кейс типа поля")
					return
				}

				if cnf.validate != nil && !cnf.validate() {
					ierr = fmt.Errorf("ошибка валидации настройки: " + fieldName)
					return
				}
			} else {
				ierr = fmt.Errorf("строку конфига нельзя присвоить: " + fieldName)
				return
			}
		}
	})

	return ierr
}
