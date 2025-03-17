package syncmap

import (
	"context"
	"kwdb/app/storage/displacement"
	"kwdb/app/storage/driver"
	"kwdb/pkg/helper"
	"sync"
	"time"
)

const DriverName = "syncmap"

type SyncMap struct {
	partitions []partition
	driver     string
	displacer  displacement.Policy
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

func (p *partition) Get(key string) (*driver.Cell, bool) {

	cell, ok := p.Get(key)

	if !ok {
		return nil, false
	}

	return cell, true
}

func (p *partition) Set(key string, cell *driver.Cell) error {
	p.vault.Store(key, cell)

	return nil
}

func (s *SyncMap) Get(ctx context.Context, key string) (*driver.Cell, error) {

	partitionIndex, pErr := helper.HashFunction(key, len(s.partitions))

	if pErr != nil {
		return nil, pErr
	}

	cell, ok := s.partitions[partitionIndex].Get(key)

	if !ok {
		return nil, nil
	}

	return cell, nil
}

func (s *SyncMap) Set(ctx context.Context, key string, value string, ttl int) error {
	cell := &driver.Cell{
		Value:   value,
		TTL:     ttl,
		AddDate: time.Now(),
	}

	partitionIndex, pErr := helper.HashFunction(key, len(s.partitions))

	if pErr != nil {
		return pErr
	}

	err := s.partitions[partitionIndex].Set(key, cell)

	if err != nil {
		return err
	}

	return nil
}

func (s *SyncMap) Delete(ctx context.Context, key string) error {

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

func (s *SyncMap) GetVaultMap() map[string]*driver.Cell {
	return make(map[string]*driver.Cell)
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
