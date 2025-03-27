package displacement

import (
	"kwdb/app"
	"kwdb/app/storage/driver"
	"kwdb/internal/helper"
	"kwdb/internal/helper/informer"
	"strconv"
	"sync"
)

type LRU struct {
	list map[string]item

	head *item
	tail *item

	lock sync.RWMutex

	memLimit uint64

	cleanerChan chan string
}

type item struct {
	prev *item
	next *item
	data *driver.Cell
	key  string
}

type remover func(key string)

func NewLRU(cchan chan string) *LRU {
	// TODO: как нормально передать, к вопросу конфигов на строне клиента
	// с методом
	// withlimit() :func()(){
	//
	// }
	memLimit := app.Config.MemLimit

	return &LRU{
		list:        make(map[string]item),
		memLimit:    memLimit,
		cleanerChan: cchan,
	}
}

func (l *LRU) Push(key string) {
	l.lock.Lock()

	ni := item{
		prev: l.head,
		next: nil,
		key:  key,
	}

	l.list[ni.key] = ni
	l.head = &ni

	l.lock.Unlock()
}

// TODO:может объединить с пушем ?
func (l *LRU) Use(key string) {

	l.lock.Lock()

	elem, ok := l.list[key]

	if ok {
		excludeConcat(&elem)
		delete(l.list, key)
	}

	l.lock.Unlock()

	l.Push(key)
}

// Cut метод очистки базы до конфигурируемого лимита
// Метод пишет в канал очистки ключи базы данных, которые необходимо удалить
// подрезает конец списка
// Если нет элементов, то no-op
func (l *LRU) Cut() {

	memAlloc := helper.AllocMB()

	informer.InfChan <- "Запущена очистка"

	for memAlloc > l.memLimit {
		l.lock.Lock()
		if len(l.list) > 0 {
			// TODO: найминг говна?
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
	}

	informer.InfChan <- strconv.FormatUint(memAlloc-helper.AllocMB(), 10) + " очищено"

	if l.tail != nil {
		l.tail.prev = nil
	}

}

func excludeConcat(i *item) {
	prev := i.prev
	next := i.next

	if prev != nil {
		prev.next = next
	}

	if next != nil {
		next.prev = prev
	}
}
