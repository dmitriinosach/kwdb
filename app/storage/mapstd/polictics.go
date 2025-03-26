package mapstd

import "kwdb/app/storage"

type Policy interface {
	Cut()
	Push(string string, cell *storage.Cell)
}
