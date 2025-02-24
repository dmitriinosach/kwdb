package driver

import (
	"context"
)

type Driver interface {
	Get(ctx context.Context, key string) (*Cell, error)
	Set(ctx context.Context, key string, value string, ttl int) error
	Delete(ctx context.Context, key string) error
	Has(ctx context.Context, key string) (bool, error)
	Info() string
	GetVaultMap() map[string]*Cell
	Truncate() bool
}
