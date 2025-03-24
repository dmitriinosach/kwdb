package displacement

import (
	"kwdb/app"
	"kwdb/app/storage/driver"
	"runtime"
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

func (l *LRU) Push(key string, cell *driver.Cell) {
	l.lock.Lock()

	ni := item{
		prev: l.head,
		next: nil,
		data: cell,
		key:  key,
	}

	l.list[ni.key] = ni
	l.head = &ni

	l.lock.Unlock()
}

// TODO:может объединить с пушем ?
func (l *LRU) Reuse(key string, cell *driver.Cell) {

	l.lock.Lock()

	elem := l.list[key]

	excludeConcat(&elem)

	delete(l.list, key)

	l.lock.Unlock()

	l.Push(key, cell)
}

// Cut метод очистки базы до конфигурируемого лимита
// Метод пишет в канал очистки ключи базы данных, которые необходимо удалить
// подрезает конец списка
// Если нет элементов , то no-op
func (l *LRU) Cut() {

	var stat runtime.MemStats

	runtime.ReadMemStats(&stat)

	memAlloc := stat.Alloc / 1024 / 1024

	if len(l.list) > 0 {
		for memAlloc > l.memLimit {
			// TODO: найминг говна?
			t := l.tail
			l.tail = l.tail.next

			delete(l.list, t.key)

			l.cleanerChan <- t.key
		}

		if l.tail != nil {
			l.tail.prev = nil
		}
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
