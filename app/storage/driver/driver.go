package driver

import (
	"context"
	"kwdb/app/storage/displacement"
	"time"
)

type Driver interface {
	Get(ctx context.Context, key string) (*Cell, error)
	Set(ctx context.Context, key string, value string, ttl int) error
	Delete(ctx context.Context, key string) error
	Info() string
	GetVaultMap() map[string]*Cell
	Truncate() bool
	SetMemPolicy(policy displacement.Policy) bool
}

type Status struct {
	Freelines string
	Started   time.Time
}
