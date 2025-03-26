// Файл содержит управление историей ввода в консоль

package main

import (
	"kwdb/internal/helper"
	"log"
	"os"
	"path/filepath"
)

const appDir = "/kwdb"
const historyDir = "/history"

type history struct {
	arr [20]string
	ptr int
}

func (h *history) Push(cmd string) {

	h.ptr = 0

	//вставка в начало
	h.arr = [20]string(append([]string{""}, h.arr[0:20]...))

	h.arr[0] = cmd

	//h.save()
}

func (h *history) Prev() string {

	if h.ptr == len(h.arr) {
		h.ptr = 0
	}

	cmd := h.arr[h.ptr]
	if cmd == "" {
		h.ptr = 0
		return ""
	}

	h.ptr++

	return cmd
}

func (h *history) save() string {
	homeDir := helper.GetUserHome()
	homeDir += appDir + historyDir
	h.write(homeDir)

	return homeDir
}

func (h *history) write(dir string) {

	// Указываем имя файла
	filename := "history.txt"

	// Формируем полный путь
	filePath := filepath.Join(dir, filename)

	// Создаем файл
	file, err := os.Open(filePath)
	if err != nil {
		err := os.Mkdir(helper.GetUserHome()+appDir, 0777)
		err = os.Mkdir(dir, 0777)
		os.Create(filePath)
		log.Printf(err.Error())
	}
	defer file.Close()

	for i := range h.arr {
		if _, err := file.Write([]byte(h.arr[i])); err != nil {
			log.Printf(err.Error())
		}
	}
}
