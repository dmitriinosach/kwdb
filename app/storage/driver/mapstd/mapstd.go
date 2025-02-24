package mapstd

import (
	"context"
	"kwdb/app/errorpkg"
	"kwdb/app/storage/driver"
	"strconv"
	"sync"
	"time"
)

const DriverName = "hashmap"

type HashMapStandard struct {
	partitions []partition
	locker     sync.RWMutex
	driver     string
}

func NewHashMapStandard() *HashMapStandard {
	return &HashMapStandard{
		partitions: make([]partition, 10),
		locker:     sync.RWMutex{},
		driver:     "hash",
	}
}

type partition struct {
	vault  map[string]*driver.Cell
	locker sync.RWMutex
}

func (p *partition) Get(key string) (*driver.Cell, bool) {
	p.locker.RLock()
	cell, ok := p.vault[key]
	p.locker.RUnlock()

	if !ok {
		return nil, false
	}

	return cell, true
}

func (p *partition) Set(key string, cell *driver.Cell) error {
	p.locker.Lock()
	p.vault[key] = cell
	p.locker.Unlock()

	return nil
}

func (s *HashMapStandard) Get(ctx context.Context, key string) (*driver.Cell, error) {

	partitionIndex, ok := driver.HashFunction(key)

	if (partitionIndex - 1) > 10 {
		return nil, errorpkg.ErrHashFunctionIndexOutRange
	}
	if !ok {
		return nil, errorpkg.ErrHashFunctionCompute
	}

	cell, ok := s.partitions[partitionIndex].Get(key)

	if !ok {
		return nil, nil
	}

	return cell, nil
}

func (s *HashMapStandard) Set(ctx context.Context, key string, value string, ttl int) error {
	cell := &driver.Cell{
		Value:   value,
		TTL:     ttl,
		AddDate: time.Now(),
	}

	partitionIndex, ok := driver.HashFunction(key)

	if !ok {
		return errorpkg.ErrHashFunctionCompute
	}
	if (partitionIndex - 1) > 10 {
		return errorpkg.ErrHashFunctionIndexOutRange
	}

	err := s.partitions[partitionIndex].Set(key, cell)

	if err != nil {
		return err
	}

	return nil
}

func (s *HashMapStandard) Delete(ctx context.Context, key string) error {

	rmin, rmax := 0, 9

	for i := rmin; i < rmax; i++ {
		s.partitions[i].locker.RLock()
		delete(s.partitions[i].vault, key)
		s.partitions[i].locker.RUnlock()
	}

	return nil
}

func (s *HashMapStandard) Has(ctx context.Context, key string) (bool, error) {

	partitionIndex, ok := driver.HashFunction(key)
	if !ok {
		return false, errorpkg.ErrHashFunctionCompute
	}
	if (partitionIndex - 1) > 10 {
		return false, errorpkg.ErrHashFunctionIndexOutRange
	}

	if s.partitions[partitionIndex].vault[key] == nil {
		return false, nil
	}

	return true, nil
}

func (s *HashMapStandard) Info() string {
	info := "driver:" + s.driver + "\n"
	info += "Length: \n"

	rmin, rmax := 0, 9
	for i := rmin; i < rmax; i++ {
		info += "partition-" + strconv.Itoa(i) + ": " + strconv.Itoa(len(s.partitions[i].vault)) + "\n"
	}

	return info
}

func (s *HashMapStandard) GetVaultMap() map[string]*driver.Cell {
	return s.partitions[0].vault
}

func (s *HashMapStandard) Truncate() bool {
	rmin, rmax := 0, 9
	for i := rmin; i < rmax; i++ {
		s.partitions[i].locker.RLock()
		clear(s.partitions[i].vault)
		s.partitions[i].locker.RUnlock()
	}

	return true
}

func (s *HashMapStandard) GetDriver() string {
	return s.driver
}
