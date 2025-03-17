package displacement

import (
	"kwdb/app/storage/driver"
	"sync"
)

type LRU struct {
	list []item
	head *item
	tail *item
	lock sync.RWMutex
}

type item struct {
	prev *item
	next *item
	data *driver.Cell
}

func NewLRU() *LRU {
	return &LRU{
		list: make([]item, 0),
	}
}

func (l *LRU) Push() {
	l.lock.Lock()

	ni := item{
		prev: l.head,
		next: nil,
		data: new(driver.Cell),
	}

	l.list = append(l.list, ni)
	l.head = &ni

	l.lock.Unlock()
}

func (l *LRU) last() *item {
	return l.tail
}

func (l *LRU) Cut() {
	l.lock.Lock()
	cnfMem := 10

	for cnfMem > 1 {
		li := l.last()

		ni := li.next

		ni.prev = nil

		l.tail = ni
	}

	l.lock.Unlock()
}
