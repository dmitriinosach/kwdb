package displacement

import (
	"kwdb/app"
	"kwdb/app/storage/driver"
	"runtime"
	"sync"
)

type LRU struct {
	list     map[string]item
	head     *item
	tail     *item
	lock     sync.RWMutex
	memLimit uint64
}

type item struct {
	prev *item
	next *item
	data *driver.Cell
	key  string
}

func NewLRU() *LRU {
	memLimit := app.Config.MemLimit

	return &LRU{
		list:     make(map[string]item),
		memLimit: memLimit,
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

// Если нет элементов no-op
func (l *LRU) Cut() {

	var stat runtime.MemStats

	runtime.ReadMemStats(&stat)

	memAlloc := int(stat.Alloc / 1024 / 1024)

	cnfMem := 100

	if len(l.list) > 0 {
		l.lock.Lock()

		for memAlloc > cnfMem {
			// TODO: найминг говна?
			t := l.tail
			l.tail = l.tail.next

			delete(l.list, t.key)

		}

		if l.tail != nil {
			l.tail.prev = nil
		}

		l.lock.Unlock()
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
