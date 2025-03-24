// Файл содержит структуру со статусом базы
package storage

import "time"

type status struct {
	Started    time.Time
	DriverName string
}

func (s *status) Uptime() time.Duration {
	return time.Since(s.Started)
}
