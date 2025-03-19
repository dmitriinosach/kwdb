package driver

import (
	"context"
	"time"
)

type Driver interface {
	Get(ctx context.Context, key string) (*Cell, error)
	Has(ctx context.Context, key string) (bool, error)
	Set(ctx context.Context, key string, value string, ttl int) error
	Delete(ctx context.Context, key string) error
	Info() string
	GetVaultMap() map[string]*Cell
	Truncate() bool
}

type Status struct {
	Freelines string
	Started   time.Time
}
