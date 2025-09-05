package driver

import (
	"kwdb/app/storage/cell"
	"sync"
)

type partitionSTD struct {
	vault  map[string]*cell.Cell
	locker sync.RWMutex
}

func NewPartitionSTD() PartitionType {
	return &partitionSTD{
		vault:  make(map[string]*cell.Cell),
		locker: sync.RWMutex{},
	}
}

func (p *partitionSTD) get(key string) (*cell.Cell, error) {
	p.locker.RLock()
	c, ok := p.vault[key]
	p.locker.RUnlock()

	if !ok {
		return nil, StorageErrNotFound
	}

	return c, nil
}

func (p *partitionSTD) set(key string, cell *cell.Cell) error {
	p.locker.Lock()
	p.vault[key] = cell
	p.locker.Unlock()

	return nil
}

func (p *partitionSTD) delete(key string) error {
	p.locker.Lock()
	delete(p.vault, key)
	p.locker.Unlock()

	return nil
}
