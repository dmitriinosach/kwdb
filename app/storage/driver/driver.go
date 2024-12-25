package driver

import (
	"context"
	"github.com/pkg/errors"
)

var ErrUnknownDriver = errors.New("unknown driver")

var ErrKeyNotFound = errors.New("key not found")

type Interface interface {
	SetValue(key string, value string, ttl int) bool
	GetValue(key string) (*Cell, bool)
	DeleteValue(key string) bool
	UpdateValue(key string) bool
	HasKey(key string) bool
	Truncate()
	Init()
	Lock()
	Unlock()
	GetDriver() string
	Info() string
	GetVaultMap() map[string]*Cell
}

type Storage interface {
	GetDriver() string
	Get(ctx context.Context, key string) (*Cell, error)
	Set(ctx context.Context, key string, value string, ttl int) error
	Delete(ctx context.Context, key string) error
	Has(ctx context.Context, key string) (bool, error)
}
