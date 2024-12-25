package driver

import (
	"context"
	"sync"
	"time"
)

type Store struct {
	vault  map[string]*Cell
	locker sync.RWMutex
	driver string
	setOps map[string]chan struct{}
}

func NewStore() *Store {
	return &Store{
		vault:  make(map[string]*Cell, 1000),
		locker: sync.RWMutex{},
		driver: "store",
		setOps: make(map[string]chan struct{}),
	}
}

func (s *Store) GetDriver() string {
	return s.driver
}

func (s *Store) Get(ctx context.Context, key string) (*Cell, error) {
	s.locker.RLock()
	defer s.locker.RUnlock()

	cell, ok := s.vault[key]
	if !ok {
		return nil, ErrKeyNotFound
	}

	return cell, nil
}

func (s *Store) Set(ctx context.Context, key string, value string, ttl int) error {
	s.locker.Lock()
	if s.setOps[key] != nil {
		close(s.setOps[key])
	}
	stopChan := make(chan struct{})
	s.setOps[key] = stopChan

	cell := Cell{
		Value:   value,
		Key:     key,
		TTL:     ttl,
		AddDate: time.Now(),
	}

	s.vault[key] = &cell
	s.locker.Unlock()

	if ttl > 0 {
		ttlDuration := time.Duration(ttl) * time.Second

		go func() {
			select {
			case <-ctx.Done():
				return
			case <-time.After(ttlDuration):
				s.locker.Lock()
				if s.setOps[key] == stopChan {
					delete(s.vault, key)
					delete(s.setOps, key)
				}
				s.locker.Unlock()
			case <-stopChan:
			}
		}()
	}

	return nil
}

func (s *Store) Delete(ctx context.Context, key string) error {
	s.locker.Lock()
	defer s.locker.Unlock()
	delete(s.vault, key)

	return nil
}

func (s *Store) Has(ctx context.Context, key string) (bool, error) {
	s.locker.RLock()
	defer s.locker.RUnlock()

	_, ok := s.vault[key]

	return ok, nil
}
