package mapstd

import (
	"context"
	"kwdb/app/helpers"
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

func NewHashMapStandard(partitionsCount int) *HashMapStandard {
	stg := &HashMapStandard{
		partitions: make([]partition, partitionsCount),
		locker:     sync.RWMutex{},
		driver:     DriverName,
	}

	for i := range stg.partitions {
		stg.partitions[i].vault = make(map[string]*driver.Cell, 100)
	}

	return stg
}

type partition struct {
	vault  map[string]*driver.Cell
	locker sync.RWMutex

	onClean bool
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

	partitionIndex, err := helpers.HashFunction(key, len(s.partitions))

	if err != nil {
		return nil, err
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

	partitionIndex, pErr := helpers.HashFunction(key, len(s.partitions))

	if pErr != nil {
		return pErr
	}

	p := &s.partitions[partitionIndex]

	err := (*p).Set(key, cell)

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

	partitionIndex, pErr := helpers.HashFunction(key, len(s.partitions))

	if pErr != nil {
		return false, pErr
	}

	if s.partitions[partitionIndex].vault[key] == nil {
		return false, nil
	}

	return true, nil
}

func (s *HashMapStandard) Info() string {
	info := "Драйвер:" + s.driver + "\n"
	info += "Инициировано секций: " + strconv.Itoa(len(s.partitions)) + " \n"

	i := 0
	for _, p := range s.partitions {
		info += "Секция-" + strconv.Itoa(i) + ": элементов- " + strconv.Itoa(len(p.vault)) + "\n"
		i++
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
