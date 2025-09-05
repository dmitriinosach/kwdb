package driver

import (
	"kwdb/app/storage/cell"
)

type PartitionType interface {
	get(key string) (*cell.Cell, error)
	set(key string, cell *cell.Cell) error
	delete(key string) error
}
