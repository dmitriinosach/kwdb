package mapstd

import (
	"context"
	"fmt"
	"kwdb/app/storage/driver"
	"kwdb/internal/helper"
	"strconv"
	"time"
)

const DriverName = "hashmap"

type HashMapStandard struct {
	partitions []partition
	driver     string
}

func NewHashMapStandard(partitionsCount int) *HashMapStandard {
	stg := &HashMapStandard{
		partitions: make([]partition, partitionsCount),
		driver:     DriverName,
	}

	for i := range stg.partitions {
		stg.partitions[i] = partition{
			vault: make(map[string]*driver.Cell),
		}
	}

	return stg
}

func (s *HashMapStandard) Get(ctx context.Context, key string) (*driver.Cell, error) {

	partitionIndex, err := helper.HashFunction(key, len(s.partitions))

	if err != nil {
		return nil, err
	}

	cell, ok := s.partitions[partitionIndex].get(key)

	if !ok {
		return nil, fmt.Errorf("ключ не найден")
	}

	return cell, nil
}

func (s *HashMapStandard) Set(ctx context.Context, key string, value string, ttl int) error {
	cell := &driver.Cell{
		Value:   value,
		TTL:     ttl,
		AddDate: time.Now(),
	}

	partitionIndex, pErr := helper.HashFunction(key, len(s.partitions))

	if pErr != nil {
		return pErr
	}

	p := &s.partitions[partitionIndex]

	err := (*p).set(key, cell)

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

func (s *HashMapStandard) Info() string {

	info := "Драйвер:" + s.driver + "\n"

	info += "Инициировано секций: " + strconv.Itoa(len(s.partitions)) + " \n"

	i := 0

	for range s.partitions {
		s.partitions[i].locker.RLock()
		info += "Секция-" + strconv.Itoa(i) + ": элементов- " + strconv.Itoa(len(s.partitions[i].vault)) + "\n"
		i++
		s.partitions[i].locker.RUnlock()
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
