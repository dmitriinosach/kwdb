// Файл содержит структуру со статусом базы
package storage

import (
	"strconv"
	"time"
)

type status struct {
	Started    time.Time
	DriverName string

	Metrics
}

func (s *status) Uptime() time.Duration {
	return time.Since(s.Started)
}

type Metrics struct {
	getHit  uint64
	getMiss uint64
	takes   uint64
}

func (m *Metrics) HitRate() string {
	if m.takes == 0 || m.getHit == 0 {
		return "N/A"
	}

	div := float64(m.getHit) / float64(m.takes) * 100

	s := strconv.FormatFloat(div, 'f', 2, 32)
	s += "%"
	return s
}

func (m *Metrics) Hit() {
	m.getHit++
	m.takes++
}

func (m *Metrics) Miss() {
	m.getMiss++
	m.takes++
}
