// Package storage
// Status реализует структуру с метриками системы и хранилища
package storage

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"strconv"
	"sync/atomic"
	"time"
)

type status struct {
	Started    time.Time
	DriverName string
	Restoring  atomic.Bool

	// Metrics смысловое отделение типа в структуру
	//для наполнения методами и отделение системы метрик от общего статуса
	Metrics
}

func (s *status) Uptime() time.Duration {
	return time.Since(s.Started)
}

type Metrics struct {
	getHit  uint64
	getMiss uint64
	takes   uint64
	size    uint64
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

var requestMetrics = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Namespace:  "api",
	Subsystem:  "http",
	Name:       "request",
	Objectives: map[float64]float64{0.5: 0.5, 0.9: 0.1, 0.99: 0.001},
}, []string{"status"})

func ObserveRequest(d time.Duration, status int) {
	requestMetrics.WithLabelValues(strconv.Itoa(status)).Observe(d.Seconds())
}
