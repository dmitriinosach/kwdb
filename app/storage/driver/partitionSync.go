package driver

import (
	"errors"
	"kwdb/app/storage/cell"
	"sync"
)

type partitionSync struct {
	vault sync.Map
}

func NewPartitionSync() PartitionType {
	return &partitionSync{
		vault: sync.Map{},
	}
}

func (p *partitionSync) get(key string) (*cell.Cell, error) {

	c, ok := p.vault.Load(key)
	if !ok {
		return nil, StorageErrNotFound
	}

	cellP, ok := c.(*cell.Cell)
	if !ok {
		return nil, errors.New("type assertion to *cell.Cell failed")
	}
	return cellP, nil
}

func (p *partitionSync) set(key string, cell *cell.Cell) error {
	p.vault.Store(key, cell)
	return nil
}

func (p *partitionSync) delete(key string) error {
	p.vault.Delete(key)

	return nil
}
