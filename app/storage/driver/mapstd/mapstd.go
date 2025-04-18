package mapstd

import (
	"fmt"
	"kwdb/app/storage/cell"
	"kwdb/app/storage/displacement"
	"kwdb/internal/helper"
	"kwdb/internal/helper/file_system"
	"strconv"
	"sync"
)

const DriverName = "hashmap"

type HashMapStandard struct {
	partitions []partition
	driver     string
	displacer  displacement.Policy
	flushMutex sync.Mutex
}

func (s *HashMapStandard) Flush() error {
	s.flushMutex.Lock()
	defer s.flushMutex.Unlock()

	backupFile, _ := file_system.ReadOrCreate("data/backup/flush.txt")

	for i := 0; i < len(s.partitions); i++ {
		for k, c := range s.partitions[i].vault {
			_, err := backupFile.WriteString(k + ":" + string(c.Value) + "\n")
			if err != nil {
				fmt.Println("error writing to backup file" + err.Error())
			}
		}
	}

	backupFile.Close()

	return nil
}

func NewHashMapStandard(partitionsCount int, policy displacement.Policy) *HashMapStandard {
	stg := &HashMapStandard{
		partitions: make([]partition, partitionsCount),
		driver:     DriverName,
	}

	for i := range stg.partitions {
		stg.partitions[i] = partition{
			vault: make(map[string]*cell.Cell),
		}
	}

	stg.displacer = policy

	return stg
}

func (s *HashMapStandard) Get(key string) (*cell.Cell, error) {

	partitionIndex, err := helper.HashFunction(key)

	if err != nil {
		return nil, err
	}

	c, ok := s.partitions[partitionIndex].get(key)

	if !ok || c.IsExpired() {
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

	s.partitions[partitionIndex].locker.Lock()
	delete(s.partitions[partitionIndex].vault, key)
	s.partitions[partitionIndex].locker.Unlock()

	return nil
}

func (s *HashMapStandard) Info() []byte {

	info := "Драйвер:" + s.driver + "\n"

	info += "Инициировано секций: " + strconv.Itoa(len(s.partitions)) + " \n"

	i := 0

	for range s.partitions {
		s.partitions[i].locker.RLock()
		info += "Секция-" + strconv.Itoa(i) + ": элементов- " + strconv.Itoa(len(s.partitions[i].vault)) + "\n"
		s.partitions[i].locker.RUnlock()
		i++
	}

	return []byte(info)
}

func (s *HashMapStandard) GetVaultMap() map[string]*cell.Cell {
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

func (s *HashMapStandard) Cleaner(cc chan string) {

	for {
		select {
		case key := <-cc:
			s.Delete(key)
		}
	}
}
