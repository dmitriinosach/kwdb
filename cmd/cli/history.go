// Файл содержит управление историей ввода в консоль

package main

type history struct {
	arr   [30]string
	count int
	ptr   int
}

func (h *history) Push(cmd string) {

	h.ptr = h.count

	if h.count == len(h.arr) {
		h.arr = [30]string(h.arr[1:30])
	} else {
		h.arr[h.count] = cmd
	}

	h.count++
}

func (h *history) Prev() string {

	if h.ptr < 0 {
		h.ptr = h.count
	}

	cmd := h.arr[h.ptr]

	h.ptr--

	return cmd
}
