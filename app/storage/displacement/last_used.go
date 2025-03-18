package displacement

import (
	"kwdb/app/storage/driver"
	"sync"
)

type LRU struct {
	list map[string]item
	head *item
	tail *item
	lock sync.RWMutex
}

type item struct {
	prev *item
	next *item
	data *driver.Cell
	key  string
}

func NewLRU() *LRU {
	return &LRU{
		list: make(map[string]item),
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

func (l *LRU) last() *item {
	return l.tail
}

func (l *LRU) Cut() {
	l.lock.Lock()
	cnfMem := 10

	for cnfMem > 1 {
		// TODO: найминг говна?
		li := l.last()

		ni := li.next

		ni.prev = nil

		l.tail = ni

		delete(l.list, li.key)

		cnfMem--
	}

	l.lock.Unlock()
}

func excludeConcat(i *item) {
	prev := i.prev
	next := i.next

	prev.next = next
	next.prev = prev
}
