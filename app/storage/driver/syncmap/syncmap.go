package syncmap

import (
	"context"
	"kwdb/app/storage/cell"
	"kwdb/app/storage/displacement"
	"kwdb/internal/helper"
	"sync"
)

const DriverName = "syncmap"

type SyncMap struct {
	partitions []partition
	driver     string
	displacer  displacement.Policy
}

func (s *SyncMap) Flush() error {
	return nil
}

func (s *SyncMap) Has(ctx context.Context, key string) (bool, error) {
	return true, nil
}

func NewSyncMap(partitionsCount int) *SyncMap {
	return &SyncMap{
		partitions: make([]partition, partitionsCount),
		driver:     DriverName,
	}
}

type partition struct {
	vault sync.Map
}

func (p *partition) Get(key string) (*cell.Cell, bool) {

	cell, ok := p.Get(key)

	if !ok {
		return nil, false
	}

	return cell, true
}

func (p *partition) Set(key string, cell *cell.Cell) error {
	p.vault.Store(key, cell)

	return nil
}

func (s *SyncMap) Get(key string) (*cell.Cell, error) {

	partitionIndex, pErr := helper.HashFunction(key)

	if pErr != nil {
		return nil, pErr
	}

	cell, ok := s.partitions[partitionIndex].Get(key)

	if !ok {
		return nil, nil
	}

	return cell, nil
}

func (s *SyncMap) Set(key string, value string, ttl int) error {
	cell := cell.NewCell(key, ttl)

	partitionIndex, pErr := helper.HashFunction(key)

	if pErr != nil {
		return pErr
	}

	err := s.partitions[partitionIndex].Set(key, cell)

	if err != nil {
		return err
	}

	return nil
}

func (s *SyncMap) Delete(key string) error {

	rmin, rmax := 0, 9

	for i := rmin; i < rmax; i++ {
		s.partitions[i].vault.Delete(key)
	}

	return nil
}

func (s *SyncMap) Info() string {
	info := "driver:" + s.driver + "\n"
	info += "Length: \n"

	rmin, rmax := 0, 9
	for i := rmin; i < rmax; i++ {
		//info += "partition-" + strconv.Itoa(i) + ": " + strconv.Itoa(len(s.partitions[i].vault)) + "\n"
	}

	return info
}

func (s *SyncMap) GetVaultMap() map[string]*cell.Cell {
	return make(map[string]*cell.Cell)
}

func (s *SyncMap) Truncate() bool {
	rmin, rmax := 0, 9
	for i := rmin; i < rmax; i++ {
		s.partitions[i].vault.Clear()
	}

	return true
}

func (s *SyncMap) GetDriver() string {
	return s.driver
}

func (s *SyncMap) SetMemPolicy(policy displacement.Policy) bool {

	s.displacer = policy

	return true
}

func (s *SyncMap) Cleaner(cc chan string) {
	for key := range cc {
		s.Delete(key)
	}
}
