package mapstd

import (
	"kwdb/app/storage/driver"
	"sync"
)

type partition struct {
	vault  map[string]*driver.Cell
	locker sync.RWMutex
}

func (p *partition) get(key string) (*driver.Cell, bool) {
	p.locker.RLock()
	cell, ok := p.vault[key]
	p.locker.RUnlock()

	if !ok {
		return nil, false
	}

	return cell, true
}

func (p *partition) set(key string, cell *driver.Cell) error {
	p.locker.Lock()
	p.vault[key] = cell
	p.locker.Unlock()

	return nil
}
