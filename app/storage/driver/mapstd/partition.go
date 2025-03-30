package mapstd

import (
	"kwdb/app/storage/cell"
	"sync"
)

type partition struct {
	vault  map[string]*cell.Cell
	locker sync.RWMutex
}

func (p *partition) get(key string) (*cell.Cell, bool) {
	p.locker.RLock()
	c, ok := p.vault[key]
	p.locker.RUnlock()

	if !ok {
		return nil, false
	}

	return c, true
}

func (p *partition) set(key string, cell *cell.Cell) error {
	p.locker.Lock()
	p.vault[key] = cell
	p.locker.Unlock()

	return nil
}
