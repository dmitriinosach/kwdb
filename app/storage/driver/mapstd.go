package driver

import (
	"errors"
	"fmt"
	"kwdb/app/storage/cell"
	"kwdb/app/storage/displacement"
	"kwdb/internal/helper"
	"strconv"
	"sync"
)

var (
	StorageErrNotFound = errors.New("not Found")
)

type HashMapStandard struct {
	partitions []PartitionType
	driverName string
	displacer  displacement.Policy
	flushMutex sync.Mutex
}

func NewHashMapStandard(partitionsCount int, policy displacement.Policy, mapType string) *HashMapStandard {
	stg := &HashMapStandard{
		partitions: make([]PartitionType, partitionsCount),
		driverName: mapType,
	}

	if mapType != "std" {
		for i := range stg.partitions {
			stg.partitions[i] = NewPartitionSTD()
		}
	}

	if mapType != "sync" {
		for i := range stg.partitions {
			stg.partitions[i] = NewPartitionSync()
		}
	}

	stg.displacer = policy

	return stg
}

func (s *HashMapStandard) Flush() error {

	return nil
}

func (s *HashMapStandard) Get(key string) (*cell.Cell, error) {

	partitionIndex, err := helper.HashFunction(key)

	if err != nil {
		return nil, err
	}

	c, ok := s.partitions[partitionIndex].get(key)

	if ok != nil || c.IsExpired() {
		return nil, fmt.Errorf("ключ не найден")
	}

	return c, nil
}

func (s *HashMapStandard) Set(key string, value []byte, ttl int) error {

	c := cell.NewCell(value, ttl)

	partitionIndex, pErr := helper.HashFunction(key)

	if pErr != nil {
		return pErr
	}

	p := &s.partitions[partitionIndex]

	err := (*p).set(key, c)

	if err != nil {
		return err
	}

	s.displacer.Push(key)

	return nil
}

func (s *HashMapStandard) Delete(key string) error {

	partitionIndex, _ := helper.HashFunction(key)

	err := s.partitions[partitionIndex].delete(key)
	if err != nil {
		return err
	}

	return nil
}

func (s *HashMapStandard) Info() []byte {

	info := "Драйвер:" + s.driverName + "\n"

	info += "Инициировано секций: " + strconv.Itoa(len(s.partitions)) + " \n"

	i := 0

	for range s.partitions {

		info += "Секция-" + strconv.Itoa(i) + ": элементов- " + "\n"

		i++
	}

	return []byte(info)
}

func (s *HashMapStandard) Truncate() bool {
	rmin, rmax := 0, 9
	for i := rmin; i < rmax; i++ {

	}

	return true
}

func (s *HashMapStandard) Cleaner(cc chan string) {

	for {
		select {
		case key := <-cc:
			s.Delete(key)
		}
	}
}
