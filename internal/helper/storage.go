package helper

import (
	"hash/fnv"
	"kwdb/app"
	"kwdb/app/errorpkg"
	"os"
)

// HashFunction хэщ функция строки в число по остатку на кол-во партиций
func HashFunction(key string) (int, error) {

	p := app.Config.Partitions

	h := fnv.New32a()
	h.Write([]byte(key))
	number := int(h.Sum32()) % p

	if number > p {
		return 0, errorpkg.ErrHashFunctionIndexOutRange
	}

	return number, nil
}

// GetUserHome Получаем домашний каталог пользователя
func GetUserHome() string {
	home := os.Getenv("HOME")
	if home == "" {
		home = os.Getenv("USERPROFILE") // для Windows
	}
	return home
}
