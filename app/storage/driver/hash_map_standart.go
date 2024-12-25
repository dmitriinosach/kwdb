package driver

import (
	"sync"
	"time"
)

type HashMapStandard struct {
	Vault  map[string]*Cell
	locker sync.RWMutex
	Driver string
	setOps map[string]chan struct{}
}

func NewHashMapStandard() *HashMapStandard {
	return &HashMapStandard{
		Vault:  make(map[string]*Cell, 1000),
		locker: sync.RWMutex{},
		Driver: "hash",
		setOps: make(map[string]chan struct{}),
	}
}

func (hashMap *HashMapStandard) Lock() {
	hashMap.locker.Lock()
}

func (hashMap *HashMapStandard) Unlock() {
	hashMap.locker.Unlock()
}

func (hashMap *HashMapStandard) Info() string {

	return ""
}

func (hashMap *HashMapStandard) GetVaultMap() map[string]*Cell {
	return hashMap.Vault
}

func (hashMap *HashMapStandard) GetDriver() string {
	return "hash"
}

func (hashMap *HashMapStandard) SetValue(key string, value string, ttl int) bool {

	cell := Cell{
		Value:   value,
		Key:     key,
		TTL:     ttl,
		AddDate: time.Now(),
	}

	hashMap.Vault[key] = &cell

	return true
}

func (hashMap *HashMapStandard) GetValue(key string) (*Cell, bool) {

	cell, ok := hashMap.Vault[key]

	if ok == false {
		return nil, false
	}

	return cell, true
}

func (hashMap *HashMapStandard) DeleteValue(key string) bool {
	delete(hashMap.Vault, key)

	return true
}

func (hashMap *HashMapStandard) UpdateValue(key string) bool {
	//TODO implement me
	panic("implement me")
}

func (hashMap *HashMapStandard) HasKey(key string) bool {
	_, ok := hashMap.Vault[key]
	if ok {
		return true
	}

	return false
}

func (hashMap *HashMapStandard) Truncate() {
	clear(hashMap.Vault)
}

func (hashMap *HashMapStandard) Init() {
	hashMap.Vault = make(map[string]*Cell, 1000)
	hashMap.locker = sync.RWMutex{}
}
