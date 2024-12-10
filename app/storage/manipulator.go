package storage

import (
	"fmt"
	"time"
)

func SetValue(key string, value string, ttl int) bool {

	cell := Cell{
		value,
		key,
		ttl,
		time.Now(),
		true,
		false,
	}

	Storage[key] = &cell

	cell.IsLocked = false

	return true
}

func GetValue(key string) *Cell {
	now := time.Now()
	cell := Storage[key]

	for {
		if !cell.IsLocked {
			break
		}
		if now.After(now.Add(1 * time.Millisecond)) {
			panic("get timeout")
		}
	}

	return Storage[key]
}

func DeleteValue(key string) bool {

	fmt.Printf("Значение удалено %v : %v", key, Storage[key])

	delete(Storage, "moo")

	return true
}

func UpdateValue(key string) bool {

	fmt.Printf("Значение обновлено %v : %v", key, Storage[key])

	delete(Storage, key)

	return true
}

func HasKey(key string) bool {

	_, ok := Storage[key]
	if ok {
		return true
	}

	return false
}
