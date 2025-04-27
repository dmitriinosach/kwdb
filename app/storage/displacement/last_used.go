package displacement

import (
	"kwdb/app"
	"kwdb/internal/helper"
	"sync"
)

type LRU struct {
	list map[string]*node

	head, tail *node

	lock sync.RWMutex

	memLimit uint64

	cleanerChan chan string
}

type node struct {
	prev *node
	next *node
	key  string
}

func NewLRU(cchan chan string) *LRU {

	memLimit := app.Config.MemLimit

	return &LRU{
		list:        make(map[string]*node),
		memLimit:    memLimit,
		cleanerChan: cchan,
	}
}

// Push Метод добавляет в начало листа элемент,
// если он уже был - склеивает лист
func (l *LRU) Push(key string) {
	l.lock.Lock()

	elem, ok := l.list[key]

	// Если элемент существует в листе, склеиваем лист
	if ok {
		if elem.next != nil {
			elem.next.prev = elem.prev
		}
		if elem.prev != nil {
			elem.prev.next = elem.next
		}
	} else {
		// иначе создаем новый элемент и добавляем в лист
		elem = &node{
			prev: l.head,
			next: nil,
			key:  key,
		}

		l.list[elem.key] = elem
	}

	// в любом случае ставим его новой головой листа
	l.head = elem

	l.lock.Unlock()
}

// Cut метод очистки базы до конфигурируемого лимита
// Метод пишет в канал очистки ключи базы данных, которые необходимо удалить
// подрезает конец списка
// Если нет элементов, то no-op
func (l *LRU) Cut() {

	memAlloc := helper.AllocMB()

	for memAlloc > l.memLimit {
		//app.InfChan <- "Запущена очистка"
		l.lock.Lock()
		if len(l.list) > 0 {

			t := l.tail
			if t != nil {
				l.tail = l.tail.next

				delete(l.list, t.key)
				l.cleanerChan <- t.key
			}

		} else {
			return
		}
		l.lock.Unlock()
		//app.InfChan <- strconv.FormatUint(memAlloc-helper.AllocMB(), 10) + " очищено"
	}

	if l.tail != nil {
		l.tail.prev = nil
	}
}
